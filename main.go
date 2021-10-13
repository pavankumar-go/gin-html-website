package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-html-website/app/router"
	"github.com/gin-html-website/database"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sigchan
		log.Println("Recieved signal to shutdown...")
		cancel()
	}()

	db := database.GetDBConnection()
	database.Migrate(db)
	router.StartServer(ctx)
}
