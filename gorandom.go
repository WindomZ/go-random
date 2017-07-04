package gorandom

import (
	"math/rand"
	"time"
)

func init() {
	// initialize the default Source to a deterministic state.
	rand.Seed(time.Now().Unix())
}
