package encrypter

import (
	"golang.org/x/crypto/bcrypt"
)

func (e *Encrypter) BCrypt(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hashedText), err
}
