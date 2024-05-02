package main

import (
	"fmt"

	"github.com/Uemerson/go-learning-gorm/user"
	"github.com/Uemerson/go-learning-gorm/user/encrypter"
	"github.com/Uemerson/go-learning-gorm/user/helper"
	"github.com/Uemerson/go-learning-gorm/user/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Failed to connect to the database.")
		return
	}

	e := encrypter.NewEncrypter()
	r := repository.NewRepository(db)
	h := helper.NewHelper()
	s := user.NewService(r, e, h)
	r.AutoMigrate()

	if _, err := s.Create("Uemerson", "uemerson@mail.com", "password"); err != nil {
		fmt.Println(err)
	}
	u, err := s.FindByEmail("uemerson@mail.com")
	if err == nil {
		fmt.Println(u)
	}
	if _, err := s.Update(&user.User{
		ID:       u.ID,
		Email:    u.Email,
		Name:     "new_name",
		Password: "new_password",
	}); err == nil {
		fmt.Println("update user")
	}
	if u, err := s.FindByEmail("uemerson@mail.com"); err == nil {
		fmt.Println(u)
	}
	if _, err := s.DeleteById(u); err == nil {
		fmt.Println("user", u.ID, "was deleted")
	}
}
