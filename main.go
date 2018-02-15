package main

import (
	"fmt"

	"github.com/NesquikMike/CoinTest/models/coins"
)

func main()  {
	fmt.Printf("It takes on average %v flips to get a tails.", coins.FlipHistory(1000))
}
