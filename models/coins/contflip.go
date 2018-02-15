package coins

var c Coin

func ContFlip() (string, int) {
	history := ""
	tailsYet := false
	numFlips := 0
	var flip string
	for tailsYet == false {
		flip = c.Flip()
		numFlips++
		history += flip
		if flip == tails {
			tailsYet = true
		}
	}
	return history, numFlips
}
