---
title: gRPC API in Python (3. Append)
parent: Tutorials
layout: default
---

# gRPC API in Python (3. Append)

## Codes

### Import

```python
import grpc
import machrpc_pb2 as mach
import machrpc_pb2_grpc as machrpc
import numpy as np 
import time
import google.protobuf.wrappers_pb2 as pb_wp
from google.protobuf.any_pb2 import Any

```

### `Any` type converters for protocol buffer

```python
def AnyString(str: str):
    pbstr = pb_wp.StringValue()
    pbstr.value = str
    anystr = Any()
    anystr.Pack(pbstr)
    return anystr

def AnyInt64(iv: int):
    pbint = pb_wp.Int64Value()
    pbint.value = iv
    anyint = Any()
    anyint.Pack(pbint)
    return anyint

def AnyFloat(fv: int):
    pbfloat = pb_wp.DoubleValue()
    pbfloat.value = fv
    anyfloat = Any()
    anyfloat.Pack(pbfloat)
    return anyfloat
```

### Generate values

```python
sample_rate = 100
start_time = 0
end_time = 1000

timeseries = np.arange(start_time, end_time, 1/sample_rate)
frequency = 3
ts = time.time_ns()

data = list[list[Any]]()
for i, t in enumerate(timeseries):
    nanot = ts + int(t*1000000000)
    value = np.sin(2 * np.pi * frequency * t)
    data.append([AnyString("python.value"), AnyInt64(nanot), AnyFloat(value)])
```

### Connect to server

```python
channel = grpc.insecure_channel('127.0.0.1:5655')
mach_stub = machbase_proto_pb2_grpc.MachbaseStub(channel)
```

### Prepare new appender

```python
appender = stub.Appender(mach.AppenderRequest(tableName="example"))
```

### Streaming writing data

```python
def ToStream(rows: list[list[Any]]):
    for row in rows:
        yield mach.AppendData(handle = appender.handle, params = row)

stub.Append(ToStream(data))
```

