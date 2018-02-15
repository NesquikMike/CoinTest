package main

import (
	"net/http"

	"github.com/gorilla/mux"

	handlerCoins "github.com/NesquikMike/CoinTest/handlers/coins"
	"github.com/NesquikMike/CoinTest/db"
)

func main()  {
	//Connects Database
	db.Init()
	defer db.Close()

	//Creates router
	r := mux.NewRouter()
	r.HandleFunc("/", handlerCoins.AvrCoin)
	http.ListenAndServe(":8080", r)
}