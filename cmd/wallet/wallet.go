package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TobiaszCudnik/demo-distributed-wallet/api/server"
	"github.com/TobiaszCudnik/demo-distributed-wallet/internal/db"
)

func main() {
	log.Println("Wallet service stating...")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	dbReady := make(chan bool)
	rpcReady := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		db.StartDb(ctx, dbReady)
		wg.Done()
	}()
	<-dbReady
	go func() {
		server.StartRpc(ctx, rpcReady)
		wg.Done()
	}()
	<-rpcReady
	log.Println("Wallet service ok")
	<-ch
	log.Println("Wallet service shutting down...")
	cancel()
	wg.Wait()
	log.Println("Bye...")
}
