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
    "create tag table EXAMPLE (name varchar(100) primary key, time datetime basetime, value double)"
```

You could delete the table when you've done with it.

```sh
machbase-neo shell sql "drop table EXAMPLE"
```

# Make waves

1. Copy code from below and save it as `wave.go`.

2. `go run wave.go`

Go code that generates sine and cosine wave data and insert them into EXAMPLE table.

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math"
    "net/http"
    "time"
)

type WriteReq struct {
    Table string       `json:"table"`
    Data  WriteReqData `json:"data"`
}

type WriteReqData struct {
    Columns []string `json:"columns"`
    Rows    [][]any  `json:"rows"`
}

func main() {
    client := http.Client{}
    // run every 500ms
    for ts := range time.Tick(500 * time.Millisecond) {
        // get values of sin, cos that have 15s period
        delta := float64(ts.UnixMilli()%15000) / 15000
        theta := 2 * math.Pi * delta
        sin, cos := math.Sin(theta), math.Cos(theta)

        // make json payload of http api
        content, _ := json.Marshal(&WriteReq{
            Table: "EXAMPLE",
            Data: WriteReqData{
                Columns: []string{"name", "time", "value"},
                Rows: [][]any{
                    {"wave.sin", ts.UTC().UnixNano(), sin},
                    {"wave.cos", ts.UTC().UnixNano(), cos},
                },
            },
        })
        rsp, err := client.Post(
            "http://127.0.0.1:5654/db/write", 
            "application/json", 
            bytes.NewBuffer(content))
        if err != nil {
            panic(err)
        }
        if rsp.StatusCode != http.StatusOK {
            panic(fmt.Errorf("response %d", rsp.StatusCode))
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