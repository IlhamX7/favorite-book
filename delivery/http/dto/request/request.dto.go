package request

import "favorite-book/domain/entity"

type RequestBookDTO struct {
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
	Rating   int    `json:"rating"`
}

func (this RequestBookDTO) ToBookEntity() entity.Book {
	return entity.Book{
		Judul:    this.Judul,
		Penerbit: this.Penerbit,
		Rating:   this.Rating,
	}
}
