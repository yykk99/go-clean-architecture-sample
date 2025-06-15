package router

import (
	"github.com/labstack/echo/v4"

	"github.com/yykk99/go-clean-architecture-sample/internal/infrastructure/database"
	"github.com/yykk99/go-clean-architecture-sample/internal/interface/controller"
	"github.com/yykk99/go-clean-architecture-sample/internal/interface/gateway/repository"
	"github.com/yykk99/go-clean-architecture-sample/internal/usecase/interactor"
)

func NewRouter(e *echo.Echo) {
	sqlHander := database.NewDatabase()
	userRepository := repository.NewUserRepository(sqlHander)
	userInteractor := interactor.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userInteractor)

	api := e.Group("/api")
	user := api.Group("/user")
	user.GET("/:id", userController.GetUserByID)
}
