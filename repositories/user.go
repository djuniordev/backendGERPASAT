package repositories

import (
	"github.com/djuniordev/SAST-MA/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func (r *UserRepository) ListarUsuarios() ([]*models.User, error) {
	var listUsers []*models.User

	if err := r.Db.Preload("Position").Find(&listUsers).Error; err != nil {
		return nil, err
	}
	return listUsers, nil
}