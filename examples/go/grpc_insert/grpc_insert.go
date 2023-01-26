package main

import (
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

	sqlText := `insert into example (name, time, value)  values (?, ?, ?)`
	if err := cli.Exec(sqlText, "n1", time.Now(), 1.234); err != nil {
		panic(err)
	}
}
