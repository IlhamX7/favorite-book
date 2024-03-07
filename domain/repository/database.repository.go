package repository

import (
	"favorite-book/domain/entity"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type DatabaseRepository interface {
	Create(value any) error
	FindAllBook(data any, conditions ...any) error
	FindOneBook(data any) (*entity.Book, error)
	Update(data any, payload any) (*entity.Book, error)
	Delete(data any) (*entity.Book, error)
	SignUp(value any) error
	SignIn(data any) (*entity.User, error)
}

type databaseRepositoryImpl struct {
	database *gorm.DB
}

func NewDatabaseRepository(database *gorm.DB) databaseRepositoryImpl {
	return databaseRepositoryImpl{
		database: database,
	}
}

func (this databaseRepositoryImpl) Create(value any) error {
	result := this.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (this databaseRepositoryImpl) FindAllBook(data any, conditions ...any) error {
	result := this.database.Find(data, conditions...)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (this databaseRepositoryImpl) FindOneBook(data any) (*entity.Book, error) {
	var book entity.Book
	result := this.database.First(&book, "id = ?", data)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return nil, result.Error
	}

	return &book, nil
}

func (this databaseRepositoryImpl) Update(data any, payload any) (*entity.Book, error) {
	var book entity.Book
	result := this.database.First(&book, "id = ?", data)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return nil, result.Error
	}

	this.database.Model(&book).Updates(payload)

	return &book, nil
}

func (this databaseRepositoryImpl) Delete(data any) (*entity.Book, error) {
	var book entity.Book
	result := this.database.Delete(&book, "id = ?", data)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return nil, result.Error
	}

	return &book, nil
}

func (this databaseRepositoryImpl) SignUp(value any) error {
	result := this.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating user:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (this databaseRepositoryImpl) SignIn(data any) (*entity.User, error) {
	var user entity.User
	result := this.database.First(&user, "username = ?", data)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error fetching user:: %v", result.Error))
		return nil, result.Error
	}

	return &user, nil
}
