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

	sqlText := `create tag table example (
		name varchar(100) primary key, 
		time datetime basetime, 
		value double
	)`
	if err := cli.Exec(sqlText); err != nil {
		panic(err)
	}
}
