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
	data := [][3]any{}
	for i := 0; i < 100; i++ {
		datum := [3]any{}
		datum[0] = fmt.Sprintf("tag-%d", i) // name
		datum[1] = time.Now()               // time
		datum[2] = 1.0001 * float64(i+1)    // value
		data = append(data, datum)
	}

	appender, err := cli.Appender(tableName)
	if err != nil {
		panic(err)
	}
	defer appender.Close()

	for _, datum := range data {
		fmt.Println("append", datum)
		err := appender.Append(datum[:]...)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("append done.")
}
