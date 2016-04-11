package random

import "math/rand"

func RandomFloat32() float32 {
	return rand.Float32()
}

func RandomFloat32n(n int32) float32 {
	return RandomFloat32() * float32(pnInt32(n))
}

func RandomFloat64() float64 {
	return rand.Float64()
}

func RandomFloat64n(n int64) float64 {
	return RandomFloat64() * float64(pnInt64(n))
}
