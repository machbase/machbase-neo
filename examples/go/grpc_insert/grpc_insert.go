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

	sqlText := `insert into example (name, time, value)  values (?, ?, ?)`
	if result := cli.Exec(sqlText, "n1", time.Now(), 1.234); result.Err() != nil {
		panic(result.Err())
	} else {
		fmt.Println(result.Message())
	}
}
