package dto

import "github.com/go-playground/validator/v10"

type PositionDto struct {
	Name string `json:"name" validate:"required"`
}

type PositionDtoResponse struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
}

func ValidaDadosDeCargo(position *PositionDto) error {
	validate := validator.New()
	if err := validate.Struct(position); err != nil {
		return err
	}
	return nil
}
