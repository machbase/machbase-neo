---
title: Go gRPC for writing
parent: Let's make waves
layout: default
order: 30
nav_order: 120
---

# Go gRPC client program writing data

gRPC version of wave generator example is also available.
Since it doesn't need to define custom data structures for payload,
relatively less lines of code are needed comparing to HTTP.

### Step 1.

Create directory `grpc_wave`

```sh
mkdir grpc_wave && cd grpc_wave
```

### Step 2.

Find [full source code from github]({{ site.examples_url }}/go/grpc_wave/grpc_wave.go)

### Step 3.

Copy source code and save it as `grpc_wave.go` or run script below in the directory.

```sh
curl -o grpc_wave.go "https://raw.githubusercontent.com/machbase/machbase/main/examples/go/grpc_wave/grpc_wave.go"
```

### Step 4.

Initialize go mod and prepare dependent modules.

```sh
go mod init wave && go mod tidy
```

### Step 5. Run

```sh
go run .
```
