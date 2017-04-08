package gorandom

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
