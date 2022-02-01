package server

import (
	"context"
	"log"
	"net"
	"os"
	"reflect"

	"github.com/TobiaszCudnik/demo-distributed-wallet/api/types"
	pb "github.com/TobiaszCudnik/demo-distributed-wallet/api/wallet_pb"
	db "github.com/TobiaszCudnik/demo-distributed-wallet/internal/db"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedWalletServiceServer
}

var rpc *grpc.Server

func (s *server) AlterBalance(ctx context.Context, in *pb.AlterBalanceReq) (*pb.AlterBalanceRes, error) {
	var id types.WalletId = types.WalletId(in.WalletId)
	w, err := db.AlterBalance(id, in.Nonce, in.Amount)
	if err != nil {
		return nil, handleErr(err)
	}
	res := &pb.AlterBalanceRes{
		WalletId: string(id),
		Balance:  w.Balance,
	}
	return res, nil
}

func (s *server) GetBalance(ctx context.Context, in *pb.GetBalanceReq) (*pb.GetBalanceRes, error) {
	var id types.WalletId = types.WalletId(in.WalletId)
	w, err := db.GetWallet(id)
	if err != nil {
		return nil, handleErr(err)
	}
	res := &pb.GetBalanceRes{
		WalletId:  string(id),
		Balance:   w.Balance,
		UpdatedAt: timestamppb.New(w.UpdatedAt),
	}
	return res, nil
}

func (s *server) CreateWallet(ctx context.Context, in *pb.CreateWalletReq) (*pb.CreateWalletRes, error) {
	w, err := db.CreateWallet(0)
	if err != nil {
		return nil, handleErr(err)
	}
	res := &pb.CreateWalletRes{
		WalletId: string(w.Id.String()),
	}
	return res, nil
}

func StartRpc(ctx context.Context, ready chan<- bool) {
	log.Println("RPC server starting...")
	listener, err := net.Listen("tcp", ":"+os.Getenv("RPC_PORT"))
	if err != nil {
		panic(err)
	}
	rpc = grpc.NewServer()
	pb.RegisterWalletServiceServer(rpc, &server{})
	go func() {
		if err = rpc.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		if err != nil {
			return
		}
		log.Println("RPC server ok")
		ready <- true
	}()
	<-ctx.Done()
	closeRcp()
}

func handleErr(err error) error {
	if reflect.TypeOf(err) == reflect.TypeOf(types.ApiError{}) {
		return err
	}
	return types.NewApiError(pb.ErrCode_INTERNAL_ERROR)
}

func closeRcp() {
	log.Println("RPC server shutting down...")
	rpc.GracefulStop()
	log.Println("RPC server stopped")
}
