package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yykk99/go-clean-architecture-sample/internal/infrastructure/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file %v", err)
	}

	router.NewRouter()
}
