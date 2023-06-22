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

	sqlText := `create tag table example (
		name varchar(100) primary key, 
		time datetime basetime, 
		value double
	)`
	if result := cli.Exec(sqlText); result.Err() != nil {
		panic(result.Err())
	} else {
		fmt.Println(result.Message())
	}
}
