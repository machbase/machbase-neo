package main

import (
	"fmt"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	cli := machrpc.NewClient()
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `drop table example`
	if result := cli.Exec(sqlText); result.Err() != nil {
		panic(result.Err())
	} else {
		fmt.Println(result.Message())
	}
}
