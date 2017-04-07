package gorandom

import "testing"

const COUNT_FLOAT int = 100

func TestRandomFloat32(t *testing.T) {
	for i := 0; i < COUNT_FLOAT; i++ {
		//t.Log(RandomFloat32())
		n := RandomFloat32()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestRandomFloat32n(t *testing.T) {
	for i := 0; i < COUNT_FLOAT; i++ {
		//t.Log(RandomFloat32n(10))
		n := RandomFloat32n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}

func TestRandomFloat64(t *testing.T) {
	for i := 0; i < COUNT_FLOAT; i++ {
		//t.Log(RandomFloat64())
		n := RandomFloat64()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestRandomFloat64n(t *testing.T) {
	for i := 0; i < COUNT_FLOAT; i++ {
		//t.Log(RandomFloat64n(10))
		n := RandomFloat64n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}
