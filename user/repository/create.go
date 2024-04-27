package repository

import "github.com/Uemerson/go-learning-gorm/user"

func (r *Repository) Create(u *user.User) (int, error) {
	result := r.db.Create(u)
	return int(result.RowsAffected), result.Error
}
