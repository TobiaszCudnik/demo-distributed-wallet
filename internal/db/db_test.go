package db

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/TobiaszCudnik/demo-distributed-wallet/api/types"
	pb "github.com/TobiaszCudnik/demo-distributed-wallet/api/wallet_pb"
	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	dbReady := make(chan bool)
	go StartDb(context.Background(), dbReady)
	<-dbReady

	exitVal := m.Run()
	closeDb()
	os.Exit(exitVal)
}

func TestConcurrency(t *testing.T) {
	w, err := CreateWallet(100)
	if err != nil {
		t.Fatal(err)
	}
	id := types.WalletId(w.Id.String())
	// request 1 (blocking)
	StubAlterDelay = 500 * time.Millisecond
	nonce := uuid.New().String()
	_, err = AlterBalance(id, nonce, 10)
	if err != nil {
		t.Fatal(err)
	}
	// request 2
	StubAlterDelay = 0
	nonce = uuid.New().String()
	w, err = AlterBalance(id, nonce, 10)
	if err != nil {
		t.Fatal(err)
	}
	if w.Balance != 120 {
		t.Fatal("Wrong balance")
	}
}

func TestConcurrencyOverdraft(t *testing.T) {
	w, err := CreateWallet(100)
	if err != nil {
		t.Fatal(err)
	}
	id := types.WalletId(w.Id.String())
	// request 1 (blocking)
	StubAlterDelay = 500 * time.Millisecond
	nonce := uuid.New().String()
	_, err = AlterBalance(id, nonce, -60)
	if err != nil {
		t.Fatal(err)
	}
	// request 2
	StubAlterDelay = 0
	nonce = uuid.New().String()
	_, err = AlterBalance(id, nonce, -60)
	if err == nil {
		t.Fatal("Overdraft")
	}
	w, err = GetWallet(id)
	if err != nil {
		t.Fatal(err)
	}
	if w.Balance != 40 {
		t.Fatal("Wrong balance")
	}

}

func TestNonce(t *testing.T) {
	nonce := uuid.New().String()
	w, err := CreateWallet(100)
	if err != nil {
		t.Fatal(err)
	}
	_, err = AlterBalance(types.WalletId(w.Id.String()), nonce, 10)
	if err != nil {
		t.Fatal(err)
	}
	_, err = AlterBalance(types.WalletId(w.Id.String()), nonce, 10)
	if err == nil {
		t.Fatal("Duplicated nonce")
	}
	if err.(types.ApiError).Code != pb.ErrCode_ALREADY_EXECUTED {
		t.Fatal("Error should be ALREADY_EXECUTED")
	}
}
