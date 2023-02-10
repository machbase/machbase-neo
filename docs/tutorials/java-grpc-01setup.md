---
title: gRPC in Java (1.Setup)
parent: Tutorials
layout: default
---

# gRPC API in Java (1.Setup)

## Java gRPC

### Make project directory

```sh
mkdir machrpc-java && cd machrpc-java
```


### Download machrpc.proto

```sh
mkdir -p src/main/proto
curl -o src/main/proto/machrpc.proto https://raw.githubusercontent.com/machbase/neo-grpc/main/proto/machrpc.proto
```

After downloading proto file, it is required to add `java_package` option in the file.

```proto
option java_package = "com.machbase.neo.rpc";
```

Next we need to generate the gRPC client from `machrpc.proto`. It is using the `protoc` with a gRPC Java plugin.

When using Gradle or Maven, the protoc build plugin can generate the necessary code as part of the build process.
Please refer to the [grpc-java READM](https://github.com/grpc/grpc-java/blob/master/README.md) 
for how to generate code from `.proto` file.

### Add dependencies in `pom.xml`

```xml
<dependencies>
    <dependency>
        <groupId>io.grpc</groupId>
        <artifactId>grpc-netty-shaded</artifactId>
        <version>1.52.1</version>
        <scope>runtime</scope>
    </dependency>
    <dependency>
        <groupId>io.grpc</groupId>
        <artifactId>grpc-protobuf</artifactId>
        <version>1.52.1</version>
    </dependency>
    <dependency>
        <groupId>io.grpc</groupId>
        <artifactId>grpc-stub</artifactId>
        <version>1.52.1</version>
    </dependency>
    <dependency> <!-- necessary for Java 9+ -->
        <groupId>org.apache.tomcat</groupId>
        <artifactId>annotations-api</artifactId>
        <version>6.0.53</version>
        <scope>provided</scope>
    </dependency>
</dependencies>
```

### Make source file

Download or paste the soruce code from the link below into `src/main/java/com/machbase/neo/example/Example.java`.

{:.note}
> Full source code of the example is available in [here](/examples/java/grpc/).

### Generate code from .proto

Maven will download all the necessary dependency and gRPC tools and generate stub codes.

```
mvn compile
```
