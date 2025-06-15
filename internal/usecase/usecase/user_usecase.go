package usecase

import (
	"errors"

	"github.com/yykk99/go-clean-architecture-sample/internal/domain/model"
)

var (
	ErrorHTTPStatusBadRequest = errors.New("ErrorHTTPStatusBadRequest")
	ErrorHTTPStatusNotFound   = errors.New("ErrorHTTPStatusNotFound")
)

type UserUsecase interface {
	GetUserByID(id int) (*model.User, error)
}
