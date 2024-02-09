package domain

import (
	"favorite-book/domain/repository"
	"favorite-book/domain/usecase"
	"favorite-book/infrastructure"
)

type Domain struct {
	BookUsecase usecase.BookUsecase
}

func ConstructDomain() Domain {
	databaseConn := infrastructure.NewDatabaseConnection()

	databaseRepository := repository.NewDatabaseRepository(databaseConn)

	bookUsecase := usecase.NewBookUsecase(databaseRepository)

	return Domain{
		BookUsecase: &bookUsecase,
	}
}
