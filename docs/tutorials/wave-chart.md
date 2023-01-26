---
parent: Tutorials
title: Wave data and monitoring
layout: default
---
# Wave data and monitoring

## Run machbase-neo server

Start machbase-neo server.

```sh
machbase-neo serve
```

## Create example table

Create `EXAMPLE` table for this course if it doesn't exist.

```sh
machbase-neo shell "create tag table EXAMPLE (name varchar(100) primary key, time datetime basetime, value double)"
```

You could delete the table when you've done with it.

```sh
machbase-neo shell "drop table EXAMPLE"
```

## Make waves

1. Find [full source code from github]({{ site.examples_url }}/go/http_wave/http_wave.go)
2. Copy source code and save it as `wave.go`.
3. `go run wave.go`

This Go code generates sine & cosine wave data and writes them into EXAMPLE table.

### Code explains

Define data structure that represents the payload of write API.

```go
type WriteReq struct {
    Table string       `json:"table"`
    Data  WriteReqData `json:"data"`
}

type WriteReqData struct {
    Columns []string `json:"columns"`
    Rows    [][]any  `json:"rows"`
}
```

The API for writing data via HTTP is explained in [here](/machbase/docs/api-http/write) 
and it expects JSON payload.

We can prepre writing payload like below code, so that write multiple records within a payload.
Assume `sin`, `cos` variables are properly initialized `float64` values.

```go
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
```

It will be encoded as JSON for writing API like below.


```json
{
    "table": "EXAMPLE",
    "data": {
        "columns":["name", "time", "value"],
        "rows": [
            [ "wave.sin", 1670380342000000000, 1.1 ],
            [ "wave.cos", 1670380343000000000, 2.2 ]
        ]
    }
}
```

Send it to server via http POST request.

```go
client := http.Client{}
rsp, err := client.Post("http://127.0.0.1:5654/db/write", 
    "application/json", bytes.NewBuffer(content))
```

Server replies `HTTP 200 OK` if it successfully wrtes data.

## Watch waves

Use machabse-neo shell, it provides simple command line tool for monitoring incoming data.

```sh
machbase-neo shell chart --range 30s EXAMPLE/wave.sin#value EXAMPLE/wave.cos#value
```

![img](chart01.jpg)

It is also possible browsing table data forward/backward with "walk" command like below.

```sh
machbase-neo shell walk "select * from EXAMPLE order by time desc"
```

![img](chart02.jpg)

{: .note }

> Machbase treats all time data in UTC as default.
> Use `--tz` option with shell command to display time in a time-zone other than 'UTC'. 
> This option accepts 'local' and tz database format (eg: 'Europe/Paris').
> 
> 
> `machbase-neo shell --tz=local walk select...`
>
> `machbase-neo shell --tz=America/Los_Angeles walk select...`