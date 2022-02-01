package server

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/TobiaszCudnik/demo-distributed-wallet/api/types"
	pb "github.com/TobiaszCudnik/demo-distributed-wallet/api/wallet_pb"
	"github.com/TobiaszCudnik/demo-distributed-wallet/internal/db"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.WalletServiceClient

func TestMain(m *testing.M) {
	dbReady := make(chan bool)
	rpcReady := make(chan bool)
	go db.StartDb(context.Background(), dbReady)
	<-dbReady
	go StartRpc(context.Background(), rpcReady)
	<-rpcReady
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(":"+os.Getenv("RPC_PORT"), opts)
	if err != nil {
		m.Run()
	}
	client = pb.NewWalletServiceClient(conn)

	exitVal := m.Run()

	conn.Close()
	os.Exit(exitVal)
}

func TestAlterBalance(t *testing.T) {
	cases := []struct {
		balance, change, expected int64
	}{
		{100, 10, 110},
		{100, -10, 90},
		{10000000, 1000000, 11000000},
		{10000000, -1000000, 9000000},
	}

	for _, test := range cases {
		name := fmt.Sprintf("%d_%d_%d", test.balance, test.change, test.expected)
		t.Run(name, createAlterBalanceTest(test.balance, test.change, test.expected))
	}
}

func createAlterBalanceTest(balance int64, change int64, expected int64) func(t *testing.T) {
	return func(t *testing.T) {
		// create
		w, err := db.CreateWallet(balance)
		if err != nil {
			t.Fatal(err)
		}
		// alter
		req := &pb.AlterBalanceReq{
			WalletId: w.Id.String(),
			Nonce:    uuid.New().String(),
			Amount:   change,
		}
		res, err := client.AlterBalance(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		// check
		w, err = db.GetWallet(types.WalletId(res.WalletId))
		if err != nil {
			t.Fatal(err)
		}

		if w.Balance != balance+change {
			t.Fatal("Wrong balance")
		}
	}
}

func TestOverdraft(t *testing.T) {
	var balance int64 = 10
	var change int64 = -50
	// create
	w, err := db.CreateWallet(balance)
	if err != nil {
		t.Fatal(err)
	}
	// alter
	req := &pb.AlterBalanceReq{
		WalletId: w.Id.String(),
		Nonce:    uuid.New().String(),
		Amount:   change,
	}
	_, err = client.AlterBalance(context.Background(), req)
	// TODO proper RCP errors
	expectedErr := pb.ErrCode_name[int32(pb.ErrCode_BALANCE_TOO_LOW)]
	if err == nil {
		t.Fatal("Overdraft")
	}
	if !strings.Contains(err.Error(), expectedErr) {
		t.Fatal(err)
	}
}

func TestGetBalance(t *testing.T) {
	var balance int64 = 100
	w, err := db.CreateWallet(balance)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.GetBalanceReq{WalletId: w.Id.String()}
	res, err := client.GetBalance(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	if res.Balance != balance {
		t.Fatal("Wrong balance")
	}
}

func TestCreateWallet(t *testing.T) {
	res, err := client.CreateWallet(context.Background(), &pb.CreateWalletReq{})
	if err != nil {
		t.Fatal(err)
	}

	w, err := db.GetWallet(types.WalletId(res.WalletId))
	if err != nil {
		t.Fatal(err)
	}
	if w == nil {
		t.Fatal("Wallet not in the DB")
	}
}
