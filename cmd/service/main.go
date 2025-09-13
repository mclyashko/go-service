package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mclyashko/go-service/internal/di"
)

func main() {
	c := di.NewContainer()

	go func() {
		c.Logger.Println("Server is running on", c.Config.Addr)
		if err := c.Server.Start(); err != nil {
			c.Logger.Println(fmt.Errorf("server error: %w", err))
		}
	}()

	stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
    <-stop

	c.Logger.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

	if err := c.Server.Stop(ctx); err != nil {
		c.Logger.Println(fmt.Errorf("shutdown error: %w", err))
	}
	c.Logger.Print("Server stopped")
}
