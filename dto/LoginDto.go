package dto

import (
	"github.com/go-playground/validator/v10"
)

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ValidaDadosDeLogin(login *LoginDto) error {
	validate := validator.New()
	if err := validate.Struct(login); err != nil {
        return err
    }
    return nil
}