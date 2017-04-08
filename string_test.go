package gorandom

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

const COUNT_STRING int = 10000

func TestString_RandomString(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomString(10)
		assert.Equal(t, len(s), 10)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}

func TestString_RandomAlphabet(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomAlphabet(10)
		assert.Equal(t, len(s), 10)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}

func TestString_RandomCaptcha(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomCaptcha(10)
		assert.Equal(t, len(s), 10)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}

func TestString_RandomNumCaptcha(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomNumCaptcha(10)
		assert.Equal(t, len(s), 10)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}

func TestString_RandomHexString(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomHexString(10)
		assert.Equal(t, len(s), 10)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}

func TestString_RandomJson(t *testing.T) {
	results := make(map[string]string)
	for i := 0; i < COUNT_STRING; i++ {
		s := RandomJson(10)
		assert.Equal(t, len(s), 201)

		_, ok := results[s]
		assert.False(t, ok)

		results[s] = s
		assert.Equal(t, len(results), i+1)
	}
}
