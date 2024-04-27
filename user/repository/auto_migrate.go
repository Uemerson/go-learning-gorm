package repository

import "github.com/Uemerson/go-learning-gorm/user"

func (r *Repository) AutoMigrate() {
	r.db.AutoMigrate(&user.User{})
	r.db.Migrator().CreateTable()
}
