package repository

import (
	"errors"

	"github.com/Uemerson/go-learning-gorm/user"
	"gorm.io/gorm"
)

func (r *Repository) FindById(id string) (*user.User, error) {
	var u user.User
	result := r.db.First(&u, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &u, result.Error
}
