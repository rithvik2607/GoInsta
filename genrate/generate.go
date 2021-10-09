package generate

import (
	"math/rand"
	"strings"
)

func GenId() string {
	// using only lowercase characters
	charSet := "abcdefghijklmnopqrstuvxyz"
	var output strings.Builder
	length := 16
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}
