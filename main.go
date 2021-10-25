package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
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

	appPath, ok := os.LookupEnv("APP_PATH")
	if !ok {
		_, file, _, _ := runtime.Caller(0)
		appPath = filepath.Dir(file)
		log.Println("APP_PATH not set, using: ", appPath)
	}

	router.StartServer(ctx, appPath)
}
