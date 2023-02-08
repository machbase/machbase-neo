---
title: gRPC API in C# (1. Setup)
parent: Tutorials
layout: default
---

# gRPC API in C# (1. Setup)

## Install dotnet-sdk

```sh
brew install dotnet-sdk
```

## Create project directory

```sh
mkdir example-csharp && cd example-csharp
```

## Create console project

```sh
dotnet new console --framework net7.0
```

## Add gRPC packages

```sh
dotnet add package Grpc.Tools
dotnet add package Grpc.Net.Client
dotnet add package Google.Protobuf
```

## Download machrpc.proto

```sh
curl -o machrpc.proto https://raw.githubusercontent.com/machbase/neo-grpc/main/proto/machrpc.proto
```

After downloading proto file, it is required to add csharp_namespace option in the file.

```proto
option csharp_namespace = "MachRpc";
```

## Add ItemGroup in `example-csharp.csproj` XML file

```xml
  <ItemGroup>
    <Protobuf Include="machrpc.proto" GrpcServices="Client"/>
  </ItemGroup>
```
