---
parent: Tutorials
title: Wave data and monitoring
layout: default
---

# Run machbase-neo server

Start machbase-neo server.

```sh
machbase-neo serve
```

# Create example table

Create `EXAMPLE` table for this course if it doesn't exist.

```sh
machbase-neo shell sql \
    "create tag table EXAMPLE (name varchar(100) primary key, , time datetime basetime, value double)"
```

You could delete the table when you've done with it.

```sh
machbase-neo shell sql "drop table EXAMPLE"
```

# Make waves

1. Make an empty directory, then copy code below and save it as `wave.go` in the direcoty.

2. `go mod init wave`

3. `go mod tidy`

4. `go run wave.go`


Go code that generates SIN and COS wave data and insert them into EXAMPLE table.

```go
package main

import (
	"math"
	"time"
	"github.com/machbase/neo-grpc/machrpc"
)

func main() {
	cli := machrpc.NewClient(machrpc.QueryTimeout(5 * time.Second))
	if err := cli.Connect("127.0.0.1:5655"); err != nil {
		panic(err)
	}
	defer cli.Disconnect()

	sqlText := `insert into example (name, time, value) values (?, ?, ?)`
	period := float64(15000)
	tick := 500
	from := time.Now()
	for ts := range time.Tick(time.Duration(tick) * time.Millisecond) {
		delta := float64(int32(ts.Sub(from).Milliseconds()))
		theta := 2 * math.Pi * math.Mod(delta, period) / period
		sin := math.Sin(theta)
		cos := math.Cos(theta)
		if err := cli.Exec(sqlText, "wave.sin", ts, sin); err != nil {
			panic(err)
		}
		if err := cli.Exec(sqlText, "wave.cos", ts, cos); err != nil {
			panic(err)
		}
	}
}
```

# Watch waves

Use machabse-neo shell for monitoring incoming data.

```sh
machbase-neo shell chart --range 30s EXAMPLE/wave.sin#value EXAMPLE/wave.cos#value
```

![img](chart01.jpg)