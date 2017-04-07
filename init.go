package gorandom

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func pnInt16(i int16) int16 {
	if i <= 0 {
		return 1
	}
	return i
}

func pnInt32(i int32) int32 {
	if i <= 0 {
		return 1
	}
	return i
}

func pnInt(i int) int {
	if i <= 0 {
		return 1
	}
	return i
}

func pnInt64(i int64) int64 {
	if i <= 0 {
		return 1
	}
	return i
}
