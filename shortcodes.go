package shortcodes

import (
	"math/rand"
	"time"
)

// List of possible shortlink characters
const charset = "0123456789" + "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Random process seed used to randomly create a shortlink code
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenShortCodeString creates a random string of characters with a specified length
func GenShortCodeString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
