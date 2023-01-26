package main

import (
	"math"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	opts := []machrpc.ClientOption{
		machrpc.QueryTimeout(5 * time.Second),
	}

	cli := machrpc.NewClient(opts...)
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `insert into example (name, time, value) values (?, ?, ?)`

	for ts := range time.Tick(500 * time.Millisecond) {
		delta := float64(ts.UnixMilli()%15000) / 15000
		theta := 2 * math.Pi * delta
		sin, cos := math.Sin(theta), math.Cos(theta)
		if err := cli.Exec(sqlText, "wave.sin", ts, sin); err != nil {
			panic(err)
		}
		if err := cli.Exec(sqlText, "wave.cos", ts, cos); err != nil {
			panic(err)
		}
	}
}
