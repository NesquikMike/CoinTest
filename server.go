package main

import (
	"net/http"

	"github.com/gorilla/mux"

	handlerCoins "github.com/NesquikMike/CoinTest/handlers/coins"
)

func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/", handlerCoins.AvrCoin)
	http.ListenAndServe(":8080", r)
}