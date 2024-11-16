package services

import (
	"github.com/djuniordev/SAST-MA/dto"
	"github.com/djuniordev/SAST-MA/models"
	"github.com/djuniordev/SAST-MA/repositories"
	"github.com/mitchellh/mapstructure"
)

type LoginService struct {
	LoginRepository *repositories.LoginRepository
}

func (s *LoginService) Execute(loginDto *dto.LoginDto) (*dto.UserDtoResponse, error) {
	loginEntity := models.NewLogin(loginDto)

	user, err := s.LoginRepository.FazerLogin(loginEntity)
	if err != nil {
		return nil, err
	}

	var userDtoResponse dto.UserDtoResponse

	if err := mapstructure.Decode(user, &userDtoResponse); err != nil {
		return nil, err
	}

	return &userDtoResponse, nil
}
