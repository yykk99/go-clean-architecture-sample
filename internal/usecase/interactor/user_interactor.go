package interactor

import (
	"github.com/yykk99/go-clean-architecture-sample/internal/domain/model"
	"github.com/yykk99/go-clean-architecture-sample/internal/domain/repository"
	"github.com/yykk99/go-clean-architecture-sample/internal/usecase/usecase"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) usecase.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u userUsecase) GetUserByID(id int) (*model.User, error) {

	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
