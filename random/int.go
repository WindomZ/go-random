package random

import "math/rand"

func RandomInt64() int64 {
	return rand.Int63()
}

func RandomInt64n(n int64) int64 {
	return rand.Int63n(pnInt64(n))
}

func RandomInt32() int32 {
	return rand.Int31()
}

func RandomInt32n(n int32) int32 {
	return rand.Int31n(pnInt32(n))
}

func RandomInt() int {
	return rand.Int()
}

func RandomIntn(n int) int {
	return rand.Intn(pnInt(n))
}

func RandomInt16() int16 {
	return RandomInt16n(32767)
}

func RandomInt16n(n int16) int16 {
	return int16(RandomIntn(int(n)))
}
