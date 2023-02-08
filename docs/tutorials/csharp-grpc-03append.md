---
title: gRPC API in C# (3. Append)
parent: Tutorials
layout: default
---

# gRPC API in C# (3. Append)

## Code

### Connect to server 

```c#
var channel = GrpcChannel.ForAddress("http://127.0.0.1:5655");
var client = new MachRpc.Machbase.MachbaseClient(channel);
```

### Prepare new appender

```c#
var appender = client.Appender(new MachRpc.AppenderRequest { TableName = "example" });
var stream = client.Append();
```

```c#
try {
    // code that use stream & appender.Handle
}
finally {
    await stream.RequestStream.CompleteAsync();
}
```

Make `Main()` as `async Task Main()` to allow awiat for async operation.

```c#
private static async Task Main(string[] args) {
    /// use await
}
```

### Write data in high speed

```c#
for (int i = 0; i < 100000; i++)
{
    var ts = new Timestamp();
    var value = 0.1234;

    long tick = TimeUtils.GetNanoseconds();
    long secs = 1_000_000_000;
    ts.Seconds = Convert.ToInt32(tick / secs);
    ts.Nanos = Convert.ToInt32(tick % secs);

    var data = new MachRpc.AppendData { Handle = appender.Handle};
    data.Params.Add(Any.Pack(new StringValue { Value = "csharp.value" }));
    data.Params.Add(Any.Pack(ts));
    data.Params.Add(Any.Pack(new DoubleValue{ Value = value }));

    await stream.RequestStream.WriteAsync(data);
}
```

### Run and count written records

```sh
dotnet run
```

```sh
machbase-neo shell "select count(*) from example where name = 'csharp.value'"
 #  COUNT(*)
─────────────
 1  100000
```


## Full source code

```csharp
using Grpc.Net.Client;
using Google.Protobuf.WellKnownTypes;
using System.Diagnostics;

internal class Program
{
    private static async Task Main(string[] args)
    {
        var channel = GrpcChannel.ForAddress("http://127.0.0.1:5655");
        var client = new MachRpc.Machbase.MachbaseClient(channel);

        // Appender example
        var appender = client.Appender(new MachRpc.AppenderRequest { TableName = "example" });
        var stream = client.Append();
        
        var stopwatch = new Stopwatch();
        stopwatch.Start();
        try
        {
            for (int i = 0; i < 100000; i++)
            {
                var ts = new Timestamp();
                var value = 0.1234;

                long tick = TimeUtils.GetNanoseconds();
                long secs = 1_000_000_000;
                ts.Seconds = Convert.ToInt32(tick / secs);
                ts.Nanos = Convert.ToInt32(tick % secs);

                var data = new MachRpc.AppendData { Handle = appender.Handle};
                data.Params.Add(Any.Pack(new StringValue { Value = "csharp.value" }));
                data.Params.Add(Any.Pack(ts));
                data.Params.Add(Any.Pack(new DoubleValue{ Value = value }));
                await stream.RequestStream.WriteAsync(data);
            }
        }
        finally
        {
            await stream.RequestStream.CompleteAsync();

            stopwatch.Stop();
            var elapsed_time = stopwatch.ElapsedMilliseconds;
            Console.WriteLine($"Elapse {elapsed_time}ms.");
        }
    }
}
```