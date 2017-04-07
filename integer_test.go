package gorandom

import "testing"

const COUNT_INT int = 100

func TestRandomInt16(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		//t.Log(RandomInt16())
		n := RandomInt16()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestRandomInt16n(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		//t.Log(RandomInt16n(10))
		n := RandomInt16n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}
