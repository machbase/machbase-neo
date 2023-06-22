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

	var count int

	sqlText := `select count(*) from example`
	row := cli.QueryRow(sqlText)
	if err := row.Scan(&count); err != nil {
		panic(err)
	}
	fmt.Println("count=", count)
}
