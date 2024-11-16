package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int      `gorm:"primary_key"`
	Name     string   `gorm:"column:nome"`
	Surname  string   `gorm:"column:sobrenome"`
	Phone    string   `gorm:"column:telefone"`
	Email    string   `gorm:"column:email"`
	Username string   `gorm:"column:usuario"`
	Password string   `gorm:"column:senha"`
	PositionID int16     `gorm:"column:cargo_id"` // ID da posição
	Position   Position `gorm:"foreignKey:PositionID"` // Definindo a chave estrangeira
}

func (User) TableName() string {
	return "tb_usuarios"
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetSurname() string {
	return u.Surname
}

func (u *User) GetPhone() string {
	return u.Phone
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetPosition() string {
	return u.Position.Name
}
