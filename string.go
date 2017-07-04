package gorandom

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

const (
	letterBytes_09aZ   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes_Base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	letterBytes_aZ     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes_09     = "0123456789"
	letterBytes_Hex    = "0123456789ABCDEF"
	letterIdxBits      = 6
	letterIdxMask      = 1<<letterIdxBits - 1
	letterIdxMax       = 63 / letterIdxBits
)

// randomString returns an string, make up the pseudo-random chars from letters.
// n is the length.
// It empty string if n <= 0.
func randomString(n int, letters string) string {
	if n <= 0 {
		return ""
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
	return string(b)
}

// RandomString returns an string, make up the pseudo-random chars from [a-zA-z].
// n is the length.
// It empty string if n <= 0.
func RandomString(n int) string {
	return randomString(n, letterBytes_aZ)
}

// RandomAlphabet returns an string, make up the pseudo-random chars from [a-zA-z].
// n is the length.
// It empty string if n <= 0.
func RandomAlphabet(n int) string {
	return randomString(n, letterBytes_aZ)
}

// RandomCaptcha returns an string, make up the pseudo-random chars from Base58.
// n is the length.
// It empty string if n <= 0.
// Recommend used in the verification code.
func RandomCaptcha(n int) string {
	return randomString(n, letterBytes_Base58)
}

// RandomNumCaptcha returns an string, make up the pseudo-random chars from [0-9].
// n is the length.
// It empty string if n <= 0.
// Recommend used in the number verification code.
func RandomNumCaptcha(n int) string {
	return randomString(n, letterBytes_09)
}

// RandomHexString returns an string, make up the pseudo-random chars from hex [0-9A-F].
// n is the length.
// It empty string if n <= 0.
func RandomHexString(n int) string {
	return randomString(n, letterBytes_Hex)
}

// RandomJson returns an JSON string, make up the pseudo-random chars from hex [a-zA-z].
// n is the size. In this JSON, key length is 6, value length is 8.
// It empty string if n <= 0.
func RandomJson(n int) string {
	m := make(map[string]string, n)
	for i := 0; i < n; i++ {
		m[RandomString(5)+strconv.Itoa(i)] = RandomString(7) + strconv.Itoa(i)
	}
	d, _ := json.Marshal(m)
	return string(d)
}
