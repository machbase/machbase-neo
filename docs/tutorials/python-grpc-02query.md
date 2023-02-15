---
title: gRPC in Python (2.Query)
parent: Tutorials
layout: default
---

# gRPC API in Python (2.Query)
{:.no_toc}

1. TOC
{:toc}

Let's make `example.py`

## Type `any` converter

The machbase-neo gRPC is relying on `"google/protobuf/any.proto` package for its data types.
It is required to define a type conversion function.

The function below is convert protobuf any type to proper python data types.


## Convert protobuf.any value to python data type

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

## Connect

### Import packages

Import gRPC runtime package and generated files.

```python
import grpc
import machrpc_pb2_grpc
import machrpc_pb2
```

### Connect to server

Make gRPC channel to server then create a machbase-neo API stub.

```python
channel = grpc.insecure_channel('127.0.0.1:5655')
mach_stub = machrpc_pb2_grpc.MachbaseStub(channel)
```

## Execute query

Run SQL query with the stub.

```python
sqlText = "select * from example order by time limit 10"
rsp = mach_stub.Query(machrpc_pb2.QueryRequest(sql=sqlText))
```

## Get columns info of result set

We can get columns meta information of result rows after executing a query.

```python
cols = mach_stub.Columns(rsp.rowsHandle)
if cols.success:
    header = ['RowNum']
    for c in cols.columns:
        header.append(f"{c.name}({c.type})  ")
    print('   '.join(header))
```

## Fetch results

Retrieve the result records by calling `Fetch`.

```python
nrow = 0
while True:
    fetch = mach_stub.RowsFetch(rsp.rowsHandle)
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
_ = mach_stub.RowsClose(rsp.rowsHandle)
```
 
{:.warning-title}
>Rows must be Closed
>
> It is important to close rows by calling `RowsClose(handle)`.
