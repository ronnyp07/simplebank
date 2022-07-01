package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt genertes a random integer between min and max
func RandonInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0-> max-min
}

//RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

//RandomOwner generates a randon owner name
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generates a randon amount of money
func RandomMoney() int64 {
	return RandonInt(0, 1000)
}

//RandomCurrency generates a randon currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]

}
