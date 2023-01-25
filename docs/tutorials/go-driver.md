---
parent: Tutorials
title: How to use Go driver
layout: default
---

# How to use Go driver

Go provides `database/sql` package you can incorporate a wide variety of databases and data access in your applications.

This tutorial gives an introductory guide to access machbase-neo with `database/sql` package in Go.

Find [full source code from github](https://github.com/machbase/machbase/blob/main/examples/go/sql_driver.go)

## Install

Add driver into your `go.mod`.

```sh
$ go get -u github.com/machbase/neo-grpc/driver
```

## Import

Add import statement of the driver.

```go
import (
    "database/sql"
    _ "github.com/machbase/neo-grpc/driver"
)
```

Let's load sql driver and connect to server.

```go
db, err := sql.Open("machbase", "127.0.0.1:5655")
if err != nil {
    panic(err)
}
defer db.Close()
```

When call `sql.Open()`, use driverName `machbase` and machbase-neo's gRPC address.

## Insert and Query

Since we get `*sql.DB` successfully,  write and read data by standard sql syntax.

```go
var tag = "tag01"
_, err = db.Exec("INSERT INTO EXAMPLE (name, time, value) VALUES(?, ?, ?)", tag, time.Now(), 1.234)
```

Use `db.QueryRow()` to count total records of same tag name.

```go
var count int
row := db.QueryRow("SELECT count(*) FROM EXAMPLE WHERE name = ?", tag)
row.Scan(&count)
```

To iterate result of query, use `db.Query()` and get `*sql.Rows`.

{: .note }
> It is import not to forget releasing `rows` as soon as possible when you finish the work to prevent leaking resource.
> 
> General pattern is using `defer rows.Close()`.

```go
rows, err := db.Query("SELECT name, time, value FROM EXAMPLE WHERE name = ? ORDER BY TIME DESC", tag)
if err != nil {
    panic(err)
}
defer rows.Close()

for rows.Next() {
    var name string
    var ts time.Time
    var value float64
    rows.Scan(&name, &ts, &value)
    fmt.Println("name:", name, "time:", ts.Local().String(), "value:", value)
}
```