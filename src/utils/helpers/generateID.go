package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateID() string {
	source := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(source)
	movieID := randomGenerator.Intn(999999) + 1
	return fmt.Sprintf("%d", movieID)
}
