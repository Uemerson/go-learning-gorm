package user

type User struct {
	ID       string `gorm:"primaryKey;size:36;"`
	Name     string `gorm:"size:255;"`
	Email    string `gorm:"size:255;index:idx_email,unique"`
	Password string `gorm:"size:72;"`
}

type Repository interface {
	Create(u *User) (int, error)
	FindByEmail(email string) (*User, error)
	AutoMigrate()
}

type Encrypter interface {
	BCrypt(text string) (string, error)
}

type Helper interface {
	UUID() string
}

type UseCase interface {
	Create(name, email, password string) (int, error)
	FindByEmail(email string) (User, error)
}
