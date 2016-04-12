package random

import "testing"

const COUNT_BYTES int = 10

func TestRandomBytes(t *testing.T) {
	for i := 0; i < COUNT_BYTES; i++ {
		t.Log(RandomBytes(10))
	}
}

func TestRandomTimeBytes(t *testing.T) {
	for i := 0; i < COUNT_BYTES; i++ {
		t.Log(RandomTimeBytes(16))
	}
}
