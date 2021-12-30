package general

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	random_string := make([]rune, length)
	for index := range random_string {
		random_string[index] = letters[rand.Intn(len(letters))]
	}
	return string(random_string)
}
