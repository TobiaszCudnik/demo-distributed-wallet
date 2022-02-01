package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/TobiaszCudnik/demo-distributed-wallet/api/types"
	pb "github.com/TobiaszCudnik/demo-distributed-wallet/api/wallet_pb"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Wallet struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	Balance   int64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type LedgerEntry struct {
	gorm.Model
	Id        uuid.UUID `gorm:"primaryKey;type:uuid"`
	WalletId  uuid.UUID `gorm:"type:uuid"`
	Wallet    Wallet    `gorm:"foreignKey:WalletId;references:Id"`
	Amount    int64
	CreatedAt time.Time
}

var db *gorm.DB
var StubAlterDelay time.Duration = 0

func StartDb(ctx context.Context, ready chan<- bool) {
	log.Println("DB connection starting...")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic("wrong DB port ")
	}
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		port)

	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic("Cant connect to the DB")
	}
	log.Println("DB connection ok")
	err = db.AutoMigrate(Wallet{}, LedgerEntry{})
	if err != nil {
		panic(err)
	}
	log.Println("DB schema ok")
	ready <- true
	<-ctx.Done()
	closeDb()
}

func AlterBalance(id types.WalletId, nonce string, amount int64) (*Wallet, error) {
	exists, err := hasNonce(nonce)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, types.NewApiError(pb.ErrCode_ALREADY_EXECUTED)
	}
	var ret *Wallet
	// create a reversabe transaction
	err = db.Transaction(func(tx *gorm.DB) error {
		// aquire the wallet's write lock (queue requests)
		w, err := getWalletLock(tx, id)
		if err != nil {
			return err
		}
		// TODO better stubs
		if StubAlterDelay > 0 {
			log.Printf("Pausing AlterBalance for %.1fs", StubAlterDelay.Seconds())
			time.Sleep(StubAlterDelay)
		}
		// check balance when spending
		if w.Balance+amount < 0 {
			return types.NewApiError(pb.ErrCode_BALANCE_TOO_LOW)
		}
		w.Balance += amount
		// log the event
		if err := createLedgerEntry(tx, id, w, nonce, amount); err != nil {
			// TODO handle PK exists as ErrCode_ALREADY_EXECUTED
			return err
		}
		// write the wallet
		res := tx.Save(w)
		if err := res.Error; err != nil {
			return err
		}
		// re-read the wallet
		res = tx.First(w, "id = ?", w.Id)
		if err := res.Error; err != nil {
			return err
		}
		// lock released, critical section left
		ret = w
		return nil
	})
	return ret, err
}

func GetWallet(id types.WalletId) (*Wallet, error) {
	uid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}
	w := &Wallet{}
	res := db.First(w, "id = ?", uid)
	if err := res.Error; err != nil {
		if err.Error() == "record not found" {
			err = types.NewApiError(pb.ErrCode_UNKNOWN_WALLET)
		}
		return nil, err
	}
	return w, nil
}

func CreateWallet(balance int64) (*Wallet, error) {
	w := &Wallet{
		Id:        uuid.New(),
		Balance:   balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := db.Create(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// private

func hasNonce(nonce string) (bool, error) {
	var count int64
	res := db.Model(&LedgerEntry{}).Where("id = ?", nonce).Count(&count)
	if err := res.Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func getWalletLock(tx *gorm.DB, id types.WalletId) (*Wallet, error) {
	uid, err := uuid.Parse(string(id))
	if err != nil {
		return nil, err
	}
	w := &Wallet{}
	result := db.Clauses(clause.Locking{Strength: "UPDATE"}).
		First(w, "id = ?", uid)
	if err := result.Error; err != nil {
		return nil, err
	}
	if w == nil {
		return nil, types.NewApiError(pb.ErrCode_UNKNOWN_WALLET)
	}
	return w, nil
}

func createLedgerEntry(tx *gorm.DB, id types.WalletId, w *Wallet, nonce string, amount int64) error {
	nonceUuid, err := uuid.Parse(nonce)
	if err != nil {
		return err
	}
	le := &LedgerEntry{
		Id:        nonceUuid,
		WalletId:  w.Id,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	if err := tx.Create(le).Error; err != nil {
		return err
	}
	return nil
}

func closeDb() {
	log.Println("DB connection closing...")
	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("DB connection closed")
	sql.Close()
	log.Println("DB connection closed")
}
