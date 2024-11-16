package services

import (
	"github.com/djuniordev/SAST-MA/dto"
	"github.com/djuniordev/SAST-MA/repositories"
	"github.com/mitchellh/mapstructure"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func (s *UserService) Execute() ([]*dto.UserDtoResponse, error) {

	listUsers, err := s.UserRepository.ListarUsuarios()
	if err != nil {
		return nil, err
	}

	userDtoResponses := make([]*dto.UserDtoResponse, 0, len(listUsers))

	for _, user := range listUsers {
		var userDtoResponse dto.UserDtoResponse
		if err := mapstructure.Decode(user, &userDtoResponse); err != nil {
			return nil, err
		}
		userDtoResponses = append(userDtoResponses, &userDtoResponse)
	}

	return userDtoResponses, nil
}
