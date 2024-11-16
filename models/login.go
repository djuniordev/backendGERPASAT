package models

import (
	"github.com/djuniordev/SAST-MA/dto"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	username string `gorm:"column:usuario"`
	password string `gorm:"column:senha"`
}

func (Login) TableName() string {
	return "tb_usuarios"
}

func NewLogin(login *dto.LoginDto) *Login {
	return &Login{
		username: login.Username,
		password: login.Password,
	}
}

func (l *Login) GetUsername() string {
	return l.username
}

func (l *Login) GetPassword() string {
	return l.password
}
