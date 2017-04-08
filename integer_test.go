package gorandom

import "testing"

const COUNT_INT int = 100

func TestInteger_RandomInt64(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt64()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestInteger_RandomInt64n(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt64n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}

func TestInteger_RandomInt32(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt32()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestInteger_RandomInt32n(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt32n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}

func TestInteger_RandomInt(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestInteger_RandomIntn(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomIntn(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}

func TestInteger_RandomInt16(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt16()
		if n < 0 {
			t.Fatal("Must be non-negative")
		}
	}
}

func TestInteger_RandomInt16n(t *testing.T) {
	for i := 0; i < COUNT_INT; i++ {
		n := RandomInt16n(10)
		if n < 0 {
			t.Fatal("Must be non-negative")
		} else if n >= 10 {
			t.Fatal("Must less than 10")
		}
	}
}
