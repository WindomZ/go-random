package gorandom

import (
	"encoding/binary"
)

// Int64ToBytes returns bytes, from int64 i to type.
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// pnInt16 returns i, as an int16, it's 1 if i less then 1.
func pnInt16(i int16) int16 {
	if i <= 0 {
		return 1
	}
	return i
}

// pnInt32 returns i, as an int32, it's 1 if i less then 1.
func pnInt32(i int32) int32 {
	if i <= 0 {
		return 1
	}
	return i
}

// pnInt returns i, as an int, it's 1 if i less then 1.
func pnInt(i int) int {
	if i <= 0 {
		return 1
	}
	return i
}

// pnInt64 returns i, as an int64, it's 1 if i less then 1.
func pnInt64(i int64) int64 {
	if i <= 0 {
		return 1
	}
	return i
}

// pnFloat32 returns i, as an float32, it's 1 if i less then 1.
func pnFloat32(i float32) float32 {
	if i <= 0 {
		return 1
	}
	return i
}

// pnFloat64 returns i, as an float64, it's 1 if i less then 1.
func pnFloat64(i float64) float64 {
	if i <= 0 {
		return 1
	}
	return i
}
