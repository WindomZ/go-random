package gorandom

import "math/rand"

// RandomInt64 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default Source.
func RandomInt64() int64 {
	return rand.Int63()
}

// RandomInt64n returns, as an int64, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomInt64n(n int64) int64 {
	return rand.Int63n(pnInt64(n))
}

// RandomInt32 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default Source.
func RandomInt32() int32 {
	return rand.Int31()
}

// RandomInt32n returns, as an int32, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomInt32n(n int32) int32 {
	return rand.Int31n(pnInt32(n))
}

// RandomInt returns a non-negative pseudo-random int from the default Source.
func RandomInt() int {
	return rand.Int()
}

// RandomIntn returns, as an int, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomIntn(n int) int {
	return rand.Intn(pnInt(n))
}

// RandomInt16 returns a non-negative pseudo-random 15-bit integer as an int16
// from the default Source.
func RandomInt16() int16 {
	return RandomInt16n(32767)
}

// RandomInt16n returns, as an int16, a non-negative pseudo-random number in [0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomInt16n(n int16) int16 {
	return int16(RandomIntn(int(n)))
}
