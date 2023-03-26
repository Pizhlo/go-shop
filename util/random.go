package util

import (
	"math/rand"
	"strings"
	"time"
	"fmt"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}


func RandomUsername() string {
	return RandomString(6)
}


func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}