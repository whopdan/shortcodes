package shortcodes

import (
	"math/rand"
	"time"
)

// List of possible shortlink characters
const charset = "0123456789" + "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Random process seed used to randomly create a shortlink code
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string of characters with a specified length using a specified set of characters
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
