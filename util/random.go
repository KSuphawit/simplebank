package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Always generate new random seed because use current time
	rand.Seed(time.Now().UnixNano())
}

// Random int between min , max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Random money between 0 - 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"THB", "EUR", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
