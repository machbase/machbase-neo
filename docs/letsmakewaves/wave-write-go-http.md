---
title: Go HTTP for writing
parent: Let's make waves
layout: default
order: 30
nav_order: 110
---

# Go HTTP client program writing data

If you are a Go programmer and prefer to write RESTful API client, this is the way to go.

### Step 1.

Find [full source code from github]({{ site.examples_url }}/go/http_wave/http_wave.go)

### Step 2.

Copy source code and save it as `http_wave.go` or just run script below

```sh
curl -o http_wave.go "https://raw.githubusercontent.com/machbase/machbase/main/examples/go/http_wave/http_wave.go"
```

### Step 3.

```sh
go run http_wave.go
```

This Go code generates sine & cosine wave data and writes them into EXAMPLE table.

## Code explains

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
and it expects to receive JSON payload.

We can prepare payload like below code, so that write multiple records within a payload.
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

Server replies `HTTP 200 OK` if it successfully writes data.
