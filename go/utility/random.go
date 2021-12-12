package utility

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateRandomInt will generate a random int between min and max.
func GenerateRandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// GenerateBookStock will generate a random int between min and max.
func GenerateBookStock(min, max int) *int {
	num := GenerateRandomInt(min, max)
	return &num
}

// GenerateRandomString will generate a string of specified size.
func GenerateRandomString(size int) string {
	alphabet := "qwertyuiopasdfghjklxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	var sb strings.Builder

	alphabetLen := len(alphabet)

	for i := 0; i < size; i++ {
		c := alphabet[rand.Intn(alphabetLen)]
		sb.WriteByte(c)
	}
	return sb.String()
}
