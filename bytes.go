package gorandom

import (
	"math/rand"
	"time"
)

var EmptyBytes []byte = []byte{}

func randomBytes(n int, letters string) []byte {
	if n <= 0 {
		return EmptyBytes
	}
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func RandomBytes(n int) []byte {
	return randomBytes(n, letterBytes_09aZ)
}

func RandomBytesByTime(n int) []byte {
	if n <= 5 {
		return EmptyBytes
	}
	b := make([]byte, n)
	bt := Int64ToBytes(time.Now().Unix())
	for i := 0; i < 5; i++ {
		b[i] = bt[i+3]
	}
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 5; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes_09aZ) {
			b[i] = letterBytes_09aZ[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
