package usecase

import (
	"favorite-book/domain/entity"
	"favorite-book/domain/repository"
	"sort"
)

type BookUsecase interface {
	SaveBook(action entity.Book) error
	GetAllBook() (*[]entity.Book, error)
}

type bookUsecaseImpl struct {
	databaseRepository repository.DatabaseRepository
}

func NewBookUsecase(databaseRepository repository.DatabaseRepository) bookUsecaseImpl {
	return bookUsecaseImpl{
		databaseRepository: databaseRepository,
	}
}

func (this *bookUsecaseImpl) SaveBook(action entity.Book) error {
	if err := this.databaseRepository.Create(&action); err != nil {
		return err
	}

	return nil
}

func (this *bookUsecaseImpl) GetAllBook() (*[]entity.Book, error) {
	var books []entity.Book
	if err := this.databaseRepository.FindAllBook(&books); err != nil {
		return nil, err
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].ID > books[j].ID
	})

	return &books, nil
}
