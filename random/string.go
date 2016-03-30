package random

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

const (
	letterBytes0  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes1  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytes2  = "0123456789"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func RandomString(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes1) {
			b[i] = letterBytes1[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func RandomCaptcha(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes2) {
			b[i] = letterBytes2[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
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
