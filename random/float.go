package random

import (
	"math/rand"
	"time"
)

func RandomFloat64() float64 {
	rand.Seed(time.Now().Unix())
	return rand.Float64()
}

func RandomFloat64n(n int64) float64 {
	rand.Seed(time.Now().Unix())
	return float64(rand.Int63n(n*100)) / 100
}
