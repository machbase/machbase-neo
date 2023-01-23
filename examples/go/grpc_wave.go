package main

import (
	"math"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func grpc_wave() {
	opts := []machrpc.ClientOption{
		machrpc.QueryTimeout(5 * time.Second),
	}

	cli := machrpc.NewClient(opts...)
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `insert into example (name, time, value) values (?, ?, ?)`

	period := float64(15000)
	tick := 500

	from := time.Now()
	for ts := range time.Tick(time.Duration(tick) * time.Millisecond) {
		delta := float64(int32(ts.Sub(from).Milliseconds()))
		theta := 2 * math.Pi * math.Mod(delta, period) / period
		sin := math.Sin(theta)
		cos := math.Cos(theta)
		if err := cli.Exec(sqlText, "wave.sin", ts, sin); err != nil {
			panic(err)
		}
		if err := cli.Exec(sqlText, "wave.cos", ts, cos); err != nil {
			panic(err)
		}
	}
}
