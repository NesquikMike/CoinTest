package requestLog

import (
	"github.com/NesquikMike/CoinTest/db"
)

type RequestLog struct {
	AverageFlips float64
}

func (r RequestLog) Log() (error) {
	_, err :=  db.Conn.Exec("INSERT INTO request_log(time_stamp, average_flips) VALUES(now(), $1)", r.AverageFlips)
	return err
}
