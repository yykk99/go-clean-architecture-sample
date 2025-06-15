package router

import "github.com/yykk99/go-clean-architecture-sample/internal/infrastructure/database"

func NewRouter() {
	database.NewDatabase()
}
