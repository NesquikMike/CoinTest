package coins

func AverageFlip(n int) (float64) {
	trials := make([]int, n)
	var ave float64
	for i := 0; i < n; i++ {
		trials[i] = ContFlip()
	}
	for i := 0; i < n; i++ {
		ave += float64(trials[i]) / float64(n)
	}
	return ave
}
