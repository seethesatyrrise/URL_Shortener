package main

import (
	"URL_Shortener/internal/app"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errCode := 0
	defer func() {
		os.Exit(errCode)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application, err := app.New(ctx)
	if err != nil {
		fmt.Println("can't init server")
		errCode = 1
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	if err = application.Run(ctx); err != nil {
		fmt.Println("can't start server")
		errCode = 1
		return
	}

	<-signalChan
	if err = application.Shutdown(ctx); err != nil {
		fmt.Println("error shutting server")
		errCode = 1
		return
	}
	fmt.Println("server stopped")
}
