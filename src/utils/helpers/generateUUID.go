package helpers

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	u, _ := uuid.NewRandom()
	userToken := u.String()
	return userToken
}
