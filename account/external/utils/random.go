package utils

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	s := rand.NewSource(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	r := rand.New(s)
	return r.Intn(n)
}
