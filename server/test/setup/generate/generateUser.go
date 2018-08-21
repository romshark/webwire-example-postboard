package generate

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rndRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStr generates a random string of n length
func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rndRunes[rand.Intn(len(rndRunes))]
	}
	return string(b)
}

// RandNum generates a random float64 number
func RandNum(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
