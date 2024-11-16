package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserDto struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required, email"`
	Cpf      string `json:"cpf" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Position string `json:"position" validate:"required"`
}

type UserDtoResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Cpf      string `json:"cpf"`
	Username string `json:"username"`
	Position PositionDtoResponse `json:"position"`
}

func ValidaDadosDeUser(user *UserDto) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
