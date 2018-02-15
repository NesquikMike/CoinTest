package coins

import (
	"net/http"
	"fmt"

	"github.com/NesquikMike/CoinTest/models/coins"
	"github.com/NesquikMike/CoinTest/models/requestLog"
)

func AvrCoin(w http.ResponseWriter, r *http.Request) {

	var x float64
	var trialHistory []string

	x, trialHistory = coins.FlipHistory(1000)
	rL := requestLog.RequestLog{x}

	if err := rL.Log(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "There is an error: %v", err)
		return
	}

	fmt.Fprintf(w, "It takes on average %v flips to get a tails. \n", rL.AverageFlips)
	for i := 0; i < len(trialHistory); i++ {
		fmt.Fprintf(w, "Trial: %v \n", trialHistory[i])
	}
}