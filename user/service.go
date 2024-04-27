package user

import (
	"errors"
)

type Service struct {
	repo    Repository
	encrypt Encrypter
	helper  Helper
}

func NewService(r Repository, e Encrypter, h Helper) *Service {
	return &Service{
		repo:    r,
		encrypt: e,
		helper:  h,
	}
}

func (s *Service) Create(name, email, password string) (int, error) {
	if u, _ := s.repo.FindByEmail(email); u != nil {
		return 0, errors.New("email is already exits")
	}
	hashedPassword, err := s.encrypt.BCrypt(password)
	if err != nil {
		return 0, err
	}
	u := User{
		ID:       s.helper.UUID(),
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
	if _, err := s.repo.Create(&u); err != nil {
		return 0, err
	}
	return 1, nil
}

func (s *Service) FindByEmail(email string) (*User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
