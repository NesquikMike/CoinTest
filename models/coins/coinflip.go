package coins

import (
	"math/rand"
	"time"
)

const tailsProb = .5

func CoinFlip() (bool) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Float64()
	if x < tailsProb {
		return true
	}
	return false
}