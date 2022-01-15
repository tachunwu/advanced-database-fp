package main

import (
	"adfp/pkg/service"
	"os"
	"os/signal"
)

func GracefulShutdown() {
	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func main() {
	svc := service.NewTxnService()
	GracefulShutdown()
	svc.Storage.Pool.Close()
}
