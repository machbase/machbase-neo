---
title: gRPC API in Python (1. Setup)
parent: Tutorials
layout: default
---

# gRPC API in Python (1. Setup)

## Python gRPC

```sh
pip3 install grpcio grpcio-tools
```

## Download `machrpc.proto` and generate code

Make a working directory.

```sh
mkdir machrpc-py && cd machrpc-py
```

Download proto file.

```sh
curl -o machrpc.proto https://raw.githubusercontent.com/machbase/neo-grpc/main/proto/machrpc.proto
```

Compile proto file into Python.

```sh
python3 -m grpc_tools.protoc \
    -I . \
    --python_out=. \
    --grpc_python_out=. \
    ./machrpc.proto
```

As result, it generates two python files `machrpc_pb2.py` and `machrpc_pb2_grpc.py`
