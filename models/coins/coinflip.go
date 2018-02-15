package coins

import (
	"math/rand"
	"time"
)

const (
	tailsProb = .5
	heads string = "H"
	tails string = "T"
)

type Coin struct {
}

func (c Coin) Flip() (string) {

	rand.Seed(time.Now().UnixNano())
	x := rand.Float64()

	if x < tailsProb {
		return heads
	} else {
		return tails
	}
}