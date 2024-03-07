package request

import "favorite-book/domain/entity"

type RequestLoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (this RequestLoginDTO) ToLoginEntity() entity.User {
	return entity.User{
		Username: this.Username,
		Password: this.Password,
	}
}
