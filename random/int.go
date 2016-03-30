package random

import (
	"math/rand"
	"time"
)

func RandomInt64() int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63()
}

func RandomInt64n(n int64) int64 {
	rand.Seed(time.Now().Unix())
	return rand.Int63n(n)
}

func RandomInt32() int32 {
	rand.Seed(time.Now().Unix())
	return rand.Int31()
}

func RandomInt32n(n int32) int32 {
	rand.Seed(time.Now().Unix())
	return rand.Int31n(n)
}

func RandomInt() int {
	rand.Seed(time.Now().Unix())
	return rand.Int()
}

func RandomIntn(n int) int {
	rand.Seed(time.Now().Unix())
	if n <= 0 {
		n = 1
	}
	return rand.Intn(n)
}

func RandomInt16() int16 {
	return int16(rand.Int())
}

func RandomInt16n(n int) int16 {
	return int16(RandomIntn(n))
}
