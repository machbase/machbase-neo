---
title: Welcome to machbase
layout: home
order: 0
nav_order: 1
has_children: true
has_toc: false
---

# Machbase Neo

✓ Easy installation - download and run <br/>
✓ Easy to learn - standard SQL and same concept of Tables, Columns with RDBMS <br/>
✓ Easy to write and query via **HTTP**, **MQTT** and **gRPC** <br/>

<span class="fs-7">
[Get Started](./docs/getting-started/){: .btn .btn-purple .mr-4 } [Download](./releases/){: .btn .btn-green }
</span>

Machbase is the fastest timeseries database in the world[^1] with minimal foot-print. It is an ideal solution for where requires scalable environment from small server that has only limited resource like Raspberry Pi to horizontally sacle-outed cluster. Machbase Neo is built on the Machbase and adds essential features that are required for IoT industry. As result Machbase Neo ingests and be queried data through various types of protocols. For example, IoT sensors can send data directly to Machbase Neo via MQTT then applications query the stored data by SQL via HTTP.


## Download & Try Machbase Neo

Find and download machbase package from [releases](./releases/) page.

## Tutorials

[Let's make waves](./docs/tutorials/wave-chart.md) tutorial is good starting point.

If you are a gopher and looking for sql/driver for Machbase Neo, please visit "[How to use Go driver](./docs/tutorials/go-driver.md)".


## Contributing

We welcome and encourage community contributions to documents and examples for other developers. Typo and broken link fixes are appreciated.

--------------

[^1]: [TPCx-IoT Performance Results](https://www.tpc.org/tpcx-iot/results/tpcxiot_perf_results5.asp?version=2)
{: .fs-1}

