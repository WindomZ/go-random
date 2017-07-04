package gorandom

import "math/rand"

// RandomFloat32 returns, as a float32, a pseudo-random number in [0.0,1.0)
// from the default Source.
func RandomFloat32() float32 {
	return rand.Float32()
}

// RandomFloat32n returns, as a float32, a pseudo-random number in [0.0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomFloat32n(n float32) float32 {
	return RandomFloat32() * pnFloat32(n)
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0)
// from the default Source.
func RandomFloat64() float64 {
	return rand.Float64()
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,n)
// from the default Source.
// It [0,1) if n <= 0.
func RandomFloat64n(n float64) float64 {
	return RandomFloat64() * pnFloat64(n)
}
