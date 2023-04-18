package main

import (
	"math"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	cli := machrpc.NewClient()
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `insert into example (name, time, value) values (?, ?, ?)`

	for ts := range time.Tick(500 * time.Millisecond) {
		delta := float64(ts.UnixMilli()%15000) / 15000
		theta := 2 * math.Pi * delta
		sin, cos := math.Sin(theta), math.Cos(theta)
		if result := cli.Exec(sqlText, "wave.sin", ts, sin); result.Err() != nil {
			panic(result.Err())
		}
		if result := cli.Exec(sqlText, "wave.cos", ts, cos); result.Err() != nil {
			panic(result.Err())
		}
	}
}
