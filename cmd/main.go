package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yykk99/go-clean-architecture-sample/internal/infrastructure/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	e := echo.New()
	router.NewRouter(e)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdonw: ", err)
	}

	log.Println("server exiting")
}
