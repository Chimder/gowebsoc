package main

import (
	"context"
	"fmt"
	"goSql/internal/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	server := server.NewServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Printf("Error listen server: %v\n", err)
		}
	}()
	fmt.Printf("listen server...")

	<-sigCh
	fmt.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")

}
