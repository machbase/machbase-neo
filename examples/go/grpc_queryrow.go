package main

import (
	"fmt"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func grpc_queryrow() {
	opts := []machrpc.ClientOption{
		machrpc.QueryTimeout(5 * time.Second),
	}

	cli := machrpc.NewClient(opts...)
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	var count int

	sqlText := `select count(*) from example`
	row := cli.QueryRow(sqlText)
	if err := row.Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println("count=", count)
}
