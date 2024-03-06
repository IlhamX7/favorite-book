package usecase

import (
	"favorite-book/domain/entity"
	"favorite-book/domain/repository"
)

type BookUsecase interface {
	SaveBook(action entity.Book) error
	GetAllBook() (*[]entity.Book, error)
	GetOneBook(param any) (*entity.Book, error)
	UpdateBook(param any, payload any) (*entity.Book, error)
	DeleteBook(param any) (*entity.Book, error)
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

	// sort.Slice(books, func(i, j int) bool {
	// 	return books[i].ID > books[j].ID
	// })

	return &books, nil
}

func (this *bookUsecaseImpl) GetOneBook(param any) (*entity.Book, error) {
	book, err := this.databaseRepository.FindOneBook(param)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (this *bookUsecaseImpl) UpdateBook(param any, payload any) (*entity.Book, error) {
	book, err := this.databaseRepository.Update(param, payload)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (this *bookUsecaseImpl) DeleteBook(param any) (*entity.Book, error) {
	book, err := this.databaseRepository.Delete(param)
	if err != nil {
		return nil, err
	}
	return book, nil
}
