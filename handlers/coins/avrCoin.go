package coins

import (
	"net/http"
	"fmt"

	"github.com/NesquikMike/CoinTest/models/coins"
)

func AvrCoin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It takes on average %v flips to get a tails.", coins.AverageFlip(1000))
}