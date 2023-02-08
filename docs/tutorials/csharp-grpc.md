---
title: How to use gRPC in C#
parent: Tutorials
layout: default
---

# How to use gRPC in C#

## Install dotnet-sdk

```
brew install dotnet-sdk
```

## Create project directory

```
mkdir example-csharp && cd example-csharp
```

## Create console project

```
dotnet new console --framework net7.0
```

## Add gRPC packages

```
dotnet add package Grpc.Tools --version 2.51.0
dotnet add package Grpc.Net.Client --version 2.51.0
dotnet add package Google.Protobuf --version 3.21.12
```

## Download machrpc.proto

```
curl -o machrpc.proto https://raw.githubusercontent.com/machbase/neo-grpc/main/proto/machrpc.proto
```

## Edit `machrpc.proto` adding namespace

```
option csharp_namespace = "MachRpc";
```

## Add ItemGroup in `.csproj` XML file

```xml
  <ItemGroup>
    <Protobuf Include="machrpc.proto" GrpcServices="Client"/>
  </ItemGroup>
```

## Code

```csharp
using Grpc.Net.Client;
using Google.Protobuf.WellKnownTypes;

internal class Program
{
    private static void Main(string[] args)
    {
        using var channel = GrpcChannel.ForAddress("http://127.0.0.1:5655");
        var client = new MachRpc.Machbase.MachbaseClient(channel);
        var resp = client.Query(new MachRpc.QueryRequest{
            Sql = "select * from example order by time limit 10"
        });

        for (int nrow = 0; true; nrow++) {
            var fetch = client.RowsFetch(resp.RowsHandle);
            if(fetch.HasNoRows ){
                break;
            }
            foreach(Any v in fetch.Values) {
                Console.WriteLine(v);
            }
        }
        client.RowsClose(resp.RowsHandle);
    }
}
```

## Run

```
dotnet run
```