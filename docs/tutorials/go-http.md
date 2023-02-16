---
layout: default
title: HTTP API in Go
parent: Tutorials
has_children: false
---

# HTTP API in Go
{: .no_toc }

1. TOC
{: toc }

## Query

### GET

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	baseURL := "http://127.0.0.1:5654"

	client := http.Client{}
	q := url.QueryEscape("select count(*) from M$SYS_TABLES where name = 'TAGDATA'")
	rsp, err := client.Get(baseURL + "/db/query?q=" + q)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	content := string(body)

	if rsp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("ERR %s %s", rsp.Status, content))
	}

	fmt.Println(content)
}
```

### POST JSON

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	addr := "http://127.0.0.1:5654/db/query"

	queryJson := `{"q":"select count(*) from M$SYS_TABLES where name = 'TAGDATA'"}`

	client := http.Client{}
	rsp, err := client.Post(addr, "application/json", bytes.NewBufferString(queryJson))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	content := string(body)

	if rsp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("ERR %s %s", rsp.Status, content))
	}

	fmt.Println(content)
}
```

### POST FormData

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	addr := "http://127.0.0.1:5654/db/query"

	data := url.Values{"q": {"select count(*) from M$SYS_TABLES where name = 'TAGDATA'"}}

	client := http.Client{}
	rsp, err := client.Post(addr, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	content := string(body)

	if rsp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("ERR %s %s", rsp.Status, content))
	}

	fmt.Println(content)
}
```

## Write

### POST JSON

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type WriteReq struct {
	Table string      `json:"table"`
	Data WriteReqData `json:"data"`
}

type WriteReqData struct {
	Columns []string `json:"columns"`
	Rows    [][]any  `json:"rows"`
}

type WriteRsp struct {
	Success bool         `json:"success"`
	Reason  string       `json:"reason"`
	Elapse  string       `json:"elapse"`
	Data    WriteRspData `json:"data"`
}

type WriteRspData struct {
	AffectedRows uint64 `json:"affectedRows"`
}

func main() {
	addr := "http://127.0.0.1:5654/db/write"

	writeReq := WriteReq{
		Table: "TAGDATA",
		Data: WriteReqData{
			Columns: []string{"name", "time", "value", "jsondata"},
			Rows: [][]any{
				{"my-car", time.Now().UnixNano(), 32.1, `{"speed":"32.1kmh","lat":37.38906,"lon":127.12182}`},
				{"my-car", time.Now().UnixNano(), 65.4, `{"speed":"65.4kmh","lat":37.38908,"lon":127.12189}`},
				{"my-car", time.Now().UnixNano(), 76.5, `{"speed":"76.5kmh","lat":37.38912,"lon":127.12195}`},
			},
		},
	}

	queryJson, _ := json.Marshal(&writeReq)
	contentType := "application/json"

	client := http.Client{}
	rsp, err := client.Post(addr, contentType, bytes.NewBuffer(queryJson))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	content := string(body)

	if rsp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("ERR %s %s", rsp.Status, content))
	}

	fmt.Println(content)
}
```