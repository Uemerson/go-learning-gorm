package repository

import (
	"github.com/Uemerson/go-learning-gorm/user"
)

func (r *Repository) DeleteById(u *user.User) (int, error) {
	result := r.db.Delete(&u)
	return int(result.RowsAffected), result.Error
}
