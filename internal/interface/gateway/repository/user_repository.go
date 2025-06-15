package repository

import (
	"github.com/yykk99/go-clean-architecture-sample/internal/domain/model"
	"github.com/yykk99/go-clean-architecture-sample/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	sqlHander *gorm.DB
}

func NewUserRepository(sqlHander *gorm.DB) repository.UserRepository {
	return &userRepository{sqlHander: sqlHander}
}

func (u *userRepository) FindByID(id int) (*model.User, error) {
	var user model.User
	if err := u.sqlHander.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.sqlHander.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Create(user *model.User) error {
	return u.sqlHander.Create(user).Error
}

func (u *userRepository) Update(user *model.User) error {
	return u.sqlHander.Save(user).Error
}

func (u *userRepository) Delete(id int) error {
	return u.sqlHander.Delete(&model.User{}, id).Error
}
