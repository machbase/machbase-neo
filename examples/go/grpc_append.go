package main

import (
	"fmt"
	"time"

	"github.com/machbase/neo-grpc/machrpc"
)

func grpc_append() {
	opts := []machrpc.ClientOption{
		machrpc.QueryTimeout(5 * time.Second),
	}

	cli := machrpc.NewClient(opts...)
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	tableName := "example"
	appender, err := cli.Appender(tableName)
	if err != nil {
		panic(err)
	}
	defer appender.Close()

	ts := time.Now()
	for i := 0; i < 100; i++ {
		datum := [3]any{}
		datum[0] = "go-append-ex"           // name
		datum[1] = ts.Add(time.Duration(i)) // time
		datum[2] = 1.0001 * float64(i+1)    // value
		err := appender.Append(datum[:]...)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("append done.")
}
