package random

import (
	"math/rand"
	"time"
)

func RandomBytes(n int) []byte {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes0) {
			b[i] = letterBytes0[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func RandomTimeBytes(n int) []byte {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	if n <= 8 {
		return b
	}
	bt := Int64ToBytes(time.Now().Unix())
	for i := 0; i < 8; i++ {
		b[i] = bt[i]
	}
	for i, cache, remain := n-9, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes0) {
			b[i] = letterBytes0[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
