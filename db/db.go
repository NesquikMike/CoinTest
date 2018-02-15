package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "michael"
	password = "max"
	dbname = "coinflipper"
)

var Conn *sql.DB

func Init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Conn, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = Conn.Ping()
	if err != nil {
		panic(err)
	}
}

func Close() {
	Conn.Close()
}
