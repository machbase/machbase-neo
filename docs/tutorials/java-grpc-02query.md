---
title: gRPC in Java (2.Query)
parent: Tutorials
layout: default
---

# gRPC API in Java (2.Query)

## Codes

### Connect to server

The channel should be closed, otherwise the threads will remain alive.

```java
ManagedChannel channel = Grpc.newChannelBuilder("127.0.0.1:5655", InsecureChannelCredentials.create()).build();
MachbaseBlockingStub stub = MachbaseGrpc.newBlockingStub(channel);
try {
    // do job
} finally {
    channel.shutdown();
}
```

### Execute Query

```java
QueryRequest.Builder builder = Machrpc.QueryRequest.newBuilder();
builder.setSql("select * from example order by time desc limit ?");
builder.addParams(Any.pack(Int32Value.of(10)));

QueryRequest req = builder.build();
QueryResponse rsp = stub.query(req);
```

### Get columns info of result set

```java
ColumnsResponse cols = stub.columns(rsp.getRowsHandle());
ArrayList<String> headers = new ArrayList<String>();
headers.add("RowNum");
for (int i = 0; i < cols.getColumnsCount(); i++) {
    Column c = cols.getColumns(i);
    headers.add(c.getName() + "(" + c.getType() + ")");
}
```

### Fetch results

```java
int nrow = 0;
RowsFetchResponse fetch = null;
while (true) {
    fetch = stub.rowsFetch(rsp.getRowsHandle());
    if (fetch == null || fetch.getHasNoRows()) {
        break;
    }
    nrow++;

    ArrayList<String> line = new ArrayList<String>();
    line.add(Integer.toString(nrow, 10));
    List<Any> row = fetch.getValuesList();
    for (Any anyv : row) {
        line.add(convpb(anyv));
    }
    System.out.println(String.join("    ", line));
}
```

### Convert `com.google.protobuf.Any`

```java
static DateTimeFormatter sdf = DateTimeFormatter.ofPattern("yyyy.MM.dd HH:mm:ss.SSS");

static String convpb(Any any) {
    try {
        switch (any.getTypeUrl()) {
            case "type.googleapis.com/google.protobuf.StringValue": {
                StringValue v = any.unpack(StringValue.class);
                return v.getValue();
            }
            case "type.googleapis.com/google.protobuf.Timestamp": {
                Timestamp v = any.unpack(Timestamp.class);
                LocalDateTime ldt = java.time.LocalDateTime.ofEpochSecond(v.getSeconds(), v.getNanos(), ZoneOffset.UTC);
                return ldt.format(sdf);
            }
            case "type.googleapis.com/google.protobuf.DoubleValue": {
                DoubleValue v = any.unpack(DoubleValue.class);
                return Double.toString(v.getValue());
            }
            default:
                return "unsupproted " + any.getTypeUrl();
        }
    } catch (Exception e) {
        return "error " + e.getMessage();
    }
}
```

### Output

```
$ mvn exec:java -Dexec.mainClass=com.machbase.neo.example.Example

[INFO] --- exec:3.1.0:java (default-cli) @ example ---
1    python.value    2023.02.09 04:38:41.919    -0.18738131458371082
2    python.value    2023.02.09 04:38:41.909    -0.36812455270521627
3    python.value    2023.02.09 04:38:41.899    -0.535826794993456
4    python.value    2023.02.09 04:38:41.889    -0.6845471059163379
5    python.value    2023.02.09 04:38:41.879    -0.8090169943791776
6    python.value    2023.02.09 04:38:41.869    -0.9048270524669701
7    python.value    2023.02.09 04:38:41.859    -0.9685831611279518
8    python.value    2023.02.09 04:38:41.849    -0.9980267284277884
9    python.value    2023.02.09 04:38:41.839    -0.9921147013124169
10    python.value    2023.02.09 04:38:41.829    -0.9510565162916061
[INFO] ------------------------------------------------------------------------
[INFO] BUILD SUCCESS
[INFO] ------------------------------------------------------------------------
```