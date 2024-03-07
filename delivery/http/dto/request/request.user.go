package request

import "favorite-book/domain/entity"

type RequestUserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (this RequestUserDTO) ToUserEntity() entity.User {
	return entity.User{
		Username: this.Username,
		Email:    this.Email,
		Password: this.Password,
		Role:     this.Role,
	}
}
