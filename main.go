package main

import (
	"fmt"
	
	"github.com/NesquikMike/CoinTest/models/coin"
)

func main()  {
	fmt.Println("It took", coin.ContFlip(), "flips to get a tails.")
}
