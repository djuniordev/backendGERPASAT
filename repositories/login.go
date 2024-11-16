package repositories

import (
	"github.com/djuniordev/SAST-MA/models"
	"gorm.io/gorm"
)

type LoginRepository struct {
	Db *gorm.DB
}

func (r *LoginRepository) FazerLogin(login *models.Login) (*models.User, error) {
	var usuario models.User

	if err := r.Db.Preload("Position").Where("usuario = ? AND senha = ?", login.GetUsername(), login.GetPassword()).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}
