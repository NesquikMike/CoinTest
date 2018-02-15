package coins

func FlipHistory(n int) (float64, []string) {
	trials := make([]int, n)
	var ave float64
	trialHistory := make([]string, n)
	for i := 0; i < n; i++ {
		trialHistory[i], trials[i] = ContFlip()
	}
	for i := 0; i < n; i++ {
		ave += float64(trials[i]) / float64(n)
	}
	return ave, trialHistory
}