package repository

import (
	"github.com/Uemerson/go-learning-gorm/user"
)

func (r *Repository) Update(u *user.User) (int, error) {
	result := r.db.Save(&u)
	return int(result.RowsAffected), result.Error
}
