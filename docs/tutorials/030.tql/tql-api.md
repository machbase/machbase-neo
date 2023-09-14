---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL API
nav_order: 01
permalink: /docs/tutorials/tql/tql-api
---

# TQL API
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-12'), 10);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-13'), 11);
> ```
>

TQL를 저장해서 HTTP 통신의 Endpoint로 사용할 수 있다.

## Output API

### CSV Format

Save this code as `output-csv.tql`

```js
SQL( `select * from example` )
CSV()
```

Run `curl` command for HTTP communication

```sh
$ curl http://127.0.0.1:5654/db/tql/output-csv.tql
```

`result`

```sh
TAG0,1628866800000000000,12
TAG0,1628953200000000000,13
```

### JSON Format

Save this code as `output-json.tql`

```js
SQL( `select * from example` )
JSON()
```

Run `curl` command for HTTP communication

```sh
$ curl http://127.0.0.1:5654/db/tql/output-json.tql
```

`result`

```sh
{
  "data": {
    "columns": [
      "NAME",
      "TIME",
      "VALUE"
    ],
    "types": [
      "string",
      "datetime",
      "double"
    ],
    "rows": [
      [
        "TAG0",
        1628866800000000000,
        12
      ],
      [
        "TAG0",
        1628953200000000000,
        13
      ]
    ]
  },
  "success": true,
  "reason": "success",
  "elapse": "770.078µs"
}
```

### Markdown Format

Save this code as `output-markdown.tql`

```js
SQL( `select * from example` )
MARKDOWN()
```

Run `curl` command for HTTP communication

```sh
$ curl http://127.0.0.1:5654/db/tql/output-markdown.tql
```

`result`

```sh
|NAME|TIME|VALUE|
|:-----|:-----|:-----|
|TAG0|1628866800000000000|12.000000|
|TAG0|1628953200000000000|13.000000|
```

## Input API

### CSV Format

Save this code as `input-csv.tql`

#### use INSERT function
{:.no_toc}

```
CSV(payload(), 
    field(0, stringType(), 'name'),
    field(1, datetimeType('ns'), 'time'),
    field(2, doubleType(), 'value'),
    header(false)
)
INSERT("name", "time", "value", table("example"))
```

#### use APPEND function
{:.no_toc}

```
CSV(payload(), 
    field(0, stringType(), 'name'),
    field(1, datetimeType('ns'), 'time'),
    field(2, doubleType(), 'value'),
    header(false)
)
APPEND(table('example'))
```

Save this text as `input-csv.csv`

```
TAG0,1628866800000000000,12
TAG0,1628953200000000000,13
```

Run `curl` command for HTTP communication

```sh
curl -X POST http://127.0.0.1:5654/db/tql/input-csv.tql \
    -H "Content-Type: application/csv" \
    --data-binary "@input-csv.csv"
```

`result`

```sh
append 2 rows (success 2, fail 0).
```

### JSON Format

JSON Format의 경우에는 `TQL Script` 를 사용해서 데이터를 정제해야한다.

```js
BYTES(payload())
SCRIPT({
  json := import("json")
  ctx := import("context")
  val := ctx.value()
  obj := json.decode(val[0])

  for i := 0; i < len(obj.data.rows); i++ {
    ctx.yieldKey(obj.data.rows[i][0], obj.data.rows[i][1], obj.data.rows[i][2])
  }
})
INSERT("name", "time", "value", table("example"))
```

`APPEND` 함수로도 입력 가능하다.

Save this text as `input-json.json`

```
{
  "data": {
    "columns": [
      "NAME",
      "TIME",
      "VALUE"
    ],
    "types": [
      "string",
      "datetime",
      "double"
    ],
    "rows": [
      [
        "TAG0",
        1628866800000000000,
        12
      ],
      [
        "TAG0",
        1628953200000000000,
        13
      ]
    ]
  },
  "success": true,
  "reason": "success",
  "elapse": "770.078µs"
}
```

Run `curl` command for HTTP communication

```sh
curl -X POST http://127.0.0.1:5654/db/tql/input-json.tql \
    -H "Content-Type: application/json" \
    --data-binary "@input-json.json"
```

`result`

```sh
2 rows inserted.
```

## Parameter

`TQL API`를 사용할 때 parameter 값을 전달할 수 있다.

### param function

`param` 함수를 사용해서 parameter 값을 가져올 수 있다.

Save this code as `param.tql`

```js
SQL( `select * from example where name = ?`, param('name'))
CSV()
```

Run `curl` command for HTTP communication

```sh
$ curl http://127.0.0.1:5654/db/tql/param.tql?name=TAG0
```

`result`

```sh
TAG0,1628694000000000000,10
TAG0,1628780400000000000,11
```

### ?? operator

`param` 함수 뒤에 `??` keyword를 사용해서 default 값을 지정할 수 있다.

Save this code as `param-default.tql`

```js
SQL( `select * from example limit ?`, param('limit') ?? 1)
CSV()
```

Run `curl` command for HTTP communication

```sh
$ curl http://127.0.0.1:5654/db/tql/param-default.tql
```

`result`

```sh
TAG0,1628694000000000000,10
```

- set parameter

```sh
$ curl http://127.0.0.1:5654/db/tql/param-default.tql?limit=2
```

`result`

```sh
TAG0,1628694000000000000,10
TAG0,1628780400000000000,11
```
