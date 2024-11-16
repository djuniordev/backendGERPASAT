package models

import (
	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Id   int16    `gorm:"primary_key;column:id"`
	Name string `gorm:"column:nome"`
}

func (Position) TableName() string {
	return "tb_cargos"
}

func NewPosition(positionName string) *Position {
	return &Position{
		Name: positionName,
	}
}

func (p *Position) GetName() string {
	return p.Name
}
