package repository

import "github.com/yykk99/go-clean-architecture-sample/internal/domain/model"

type UserRepository interface {
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}
