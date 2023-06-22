package main

import (
	"fmt"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	cli := machrpc.NewClient()
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	var limit = 3

	sqlText := `select name, time, value from example limit ?`
	rows, err := cli.Query(sqlText, limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var ts time.Time
		var value float64
		err = rows.Scan(&name, &ts, &value)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, ts, value)
	}
}
