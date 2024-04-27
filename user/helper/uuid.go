package helper

import (
	"github.com/google/uuid"
)

func (h *Helper) UUID() string {
	return uuid.New().String()
}
