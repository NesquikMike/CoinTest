package coin

func ContFlip() int {
	tailsYet := false
	numFlips := 0
	for tailsYet == false {
		tailsYet = CoinFlip()
		numFlips++
	}
	return numFlips
}
