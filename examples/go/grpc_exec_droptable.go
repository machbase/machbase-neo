package main

import (
	"fmt"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	fmt.Println("exec example")

	opts := []machrpc.ClientOption{
		machrpc.QueryTimeout(5 * time.Second),
	}

	cli := machrpc.NewClient(opts...)
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `drop table example`
	if err := cli.Exec(sqlText); err != nil {
		panic(err)
	}
}
