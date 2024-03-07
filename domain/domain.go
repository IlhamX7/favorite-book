package domain

import (
	"favorite-book/domain/repository"
	"favorite-book/domain/usecase"
	"favorite-book/infrastructure"
)

type Domain struct {
	BookUsecase usecase.BookUsecase
	UserUsecase usecase.UserUsecase
}

func ConstructDomain() Domain {
	databaseConn := infrastructure.NewDatabaseConnection()

	databaseRepository := repository.NewDatabaseRepository(databaseConn)

	bookUsecase := usecase.NewBookUsecase(databaseRepository)

	userUsecase := usecase.NewUserUsecase(databaseRepository)

	return Domain{
		BookUsecase: &bookUsecase,
		UserUsecase: &userUsecase,
	}
}
