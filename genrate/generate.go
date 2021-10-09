package generate

import (
	"math/rand"
	"strings"
	"time"
)

/*
GenId - generates a random string from lowercase
alphabets and numbers seeding is done using time
to create a unique ID each time
*/
func GenId() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	// Set the size of ID to 8
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	// Convert string builder object to string
	str := b.String()
	return str
}
