package main

import (
	"context"
	"fmt"
	_ "goSql/docs"
	"goSql/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//		@title			Manka Api
//		@version		1.0
//		@description	Manga search
//	 @BasePath	/
func main() {
	srv := server.NewServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error listen server: %v\n", err)
		}
	}()
	fmt.Printf("Server is listening on port %s...\n", srv.Addr())

	<-sigCh
	fmt.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	srv.Close()
	fmt.Println("Server gracefully stopped")
}
