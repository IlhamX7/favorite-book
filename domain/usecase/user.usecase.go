package usecase

import (
	"favorite-book/domain/entity"
	"favorite-book/domain/repository"
)

type UserUsecase interface {
	SaveUser(action entity.User) error
	Login(param any) (*entity.User, error)
}

type userUsecaseImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewUserUsecase(databaseRepository repository.DatabaseRepository) userUsecaseImpl {
	return userUsecaseImpl{
		databaseRepository: databaseRepository,
	}
}

func (this *userUsecaseImpl) SaveUser(action entity.User) error {
	if err := this.databaseRepository.SignUp(&action); err != nil {
		return err
	}

	return nil
}

func (this *userUsecaseImpl) Login(param any) (*entity.User, error) {
	user, err := this.databaseRepository.SignIn(param)
	if err != nil {
		return nil, err
	}
	return user, nil
}
