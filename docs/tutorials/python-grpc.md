---
title: How to use gRPC in Python
parent: Tutorials
layout: default
---

# How to use gRPC in Python

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

## Codes

Let's make `example.py`

### import modules

```python
import grpc
import machrpc_pb2
import machrpc_pb2_grpc
```

### Connect to server

```python
channel = grpc.insecure_channel('127.0.0.1:5655')
mach_stub = machbase_proto_pb2_grpc.MachbaseStub(channel)
```

### Execute query

```python
sqlText = "select * from example order by time limit 10"
rsp = stub.Query(mach.QueryRequest(sql=sqlText))
```

### Get columns info of result set

```python
cols = stub.Columns(rsp.rowsHandle)
if cols.success:
    header = ['RowNum']
    for c in cols.columns:
        header.append(f"{c.name}({c.type})  ")
    print('   '.join(header))
```

### Fetch results

```python
nrow = 0
while True:
    fetch = stub.RowsFetch(rsp.rowsHandle)
    if fetch.hasNoRows:
        break
    nrow+=1
    line = []
    line.append(str(nrow))
    for i, c in enumerate(cols.columns):
        v = fetch.values[i]
        if c.type == "string":
            line.append(convpb(v))
        elif c.type == "datetime":
            line.append(convpb(v))
        elif c.type == "double":
            line.append(convpb(v))
        else:
            line.append(f"unknown {str(v)}")
    print('     '.join(line))
_ = stub.RowsClose(rsp.rowsHandle)
```
 
{:.warning-title}
>Close rows
>
> It is important to close rows by calling `RowsClose(handle)`.

### Convert protobuf.any value to python data type

```python
from google.protobuf.any_pb2 import Any
import google.protobuf.timestamp_pb2 as pb_ts
import google.protobuf.wrappers_pb2 as pb_wp
import time
from datetime import datetime

def convpb(v):
    if v.type_url == "type.googleapis.com/google.protobuf.StringValue":
        r = pb_wp.StringValue()
        v.Unpack(r)
        return r.value
    elif v.type_url == "type.googleapis.com/google.protobuf.Timestamp":
        r = pb_ts.Timestamp()
        v.Unpack(r)
        dt = datetime.fromtimestamp(r.seconds)
        return dt.strftime('%Y-%m-%d %H:%M:%S')
    elif v.type_url == "type.googleapis.com/google.protobuf.DoubleValue":
        r = pb_wp.DoubleValue()
        v.Unpack(r)
        return str(r.value)
```