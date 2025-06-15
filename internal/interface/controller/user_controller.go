package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yykk99/go-clean-architecture-sample/internal/usecase/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *UserController {
	return &UserController{userUsecase: userUsecase}
}

func (u *UserController) GetUserByID(c echo.Context) error {
	_id := c.Param("id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, usecase.ErrorHTTPStatusBadRequest)
	}
	user, err := u.userUsecase.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, usecase.ErrorHTTPStatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}
