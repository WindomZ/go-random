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

func RandomString(n int) string {
	return randomString(n, letterBytes_aZ)
}

func RandomAlphabet(n int) string {
	return randomString(n, letterBytes_aZ)
}

func RandomCaptcha(n int) string {
	return randomString(n, letterBytes_Base58)
}

func RandomNumCaptcha(n int) string {
	return randomString(n, letterBytes_09)
}

func RandomHexString(n int) string {
	return randomString(n, letterBytes_Hex)
}

func RandomJson(n int) string {
	m := make(map[string]string, n)
	for i := 0; i < n; i++ {
		m[RandomString(5)+strconv.Itoa(i)] = RandomString(7) + strconv.Itoa(i)
	}
	d, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(d)
}
