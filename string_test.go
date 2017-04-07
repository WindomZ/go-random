package gorandom

import "testing"

const COUNT_STRING int = 10

func TestRandomString(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomString(10))
	}
}

func TestRandomAlphabet(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomAlphabet(10))
	}
}

func TestRandomCaptcha(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomCaptcha(10))
	}
}

func TestRandomNumCaptcha(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomNumCaptcha(10))
	}
}

func TestRandomHexString(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomHexString(10))
	}
}

func TestRandomJson(t *testing.T) {
	for i := 0; i < COUNT_STRING; i++ {
		t.Log(RandomJson(10))
	}
}
