---
title: Welcome to machbase
layout: home
order: 0
nav_order: 1
has_children: true
has_toc: false
---

# Machbase Neo

✓ Performant timeseries database <br/>
✓ Scalable from Raspberry Pi to high-end servers <br/>
✓ Easy to install - instant download and run <br/>
✓ Easy to learn - familiar SQL with Tables and Columns <br/>
✓ Easy to write and query via **HTTP**, **MQTT** and **gRPC** <br/>



<span class="fs-6">
[Get Started](./docs/getting-started/){: .btn .btn-purple .mr-4 } [Download](./releases/){: .btn .btn-green } 
</span>
{: .d-inline-block .v-align-top}

Latest *{{site.latest_version}}*
{:.label .label-green }

Machbase is the fastest timeseries database in the world[^1] with minimal foot-print. It is an ideal solution for where requires scalable environment from small server that has only limited resource like Raspberry Pi to horizontally scale-outed cluster. Machbase Neo is built on the Machbase and adds essential features that are required for IoT industry. Machbase Neo ingests and handls query data through various protocols. For example, IoT sensors can send data directly to Machbase Neo via MQTT then applications query the stored data with SQL via HTTP.

## Download 

### Instant download

Paste the script below in shell prompt for the latest version of recommended edition.

```sh
sh -c "$(curl -fsSL https://neo.machbase.com/assets/install.sh)" sh {{site.latest_version}}
```

### Choose edition and version manually

Find and download file of proper edition and platform from [releases](./releases/) page.

## Tutorials

[Let's make waves](./docs/tutorials/wave-chart.md) tutorial will be a good starting point.

If you are a gopher and looking for sql/driver for Machbase Neo, please visit "[How to use Go driver](./docs/tutorials/go-driver.md)".


## Contributing

We welcome and encourage community contributions to documents and examples for other developers. Typo and broken link fixes are appreciated.

--------------

[^1]: [TPCx-IoT Performance Results](https://www.tpc.org/tpcx-iot/results/tpcxiot_perf_results5.asp?version=2)
{: .fs-1}

