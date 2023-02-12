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

Machbase is the world's fastest timeseries database[^1] with a minimal footprint. It's an ideal solution for environments that require scalability, from small servers with limited resources like the Raspberry Pi to horizontally scaled clusters. Machbase Neo, built on Machbase, adds crucial features required for the IoT industry. It ingests and handles query data through various protocols, such as MQTT for direct data transfer from IoT sensors, and SQL via HTTP for data retrieval by applications.

## Download 

### Instant download

Paste the script below into the shell prompt for the latest version of the recommended edition.

```sh
sh -c "$(curl -fsSL https://neo.machbase.com/assets/install.sh)"
```

### Choose edition and version manually

Find and download file of the proper edition and platform from the [releases](./releases/) page.

## Tutorials

[Let's make waves](./docs/letsmakewaves/00.index.md) tutorial will be a good starting point.

[Raspberry PI as IoT server](./docs/tutorials/raspi-server.md) shows how to install machbase-neo on ![raspi](/assets/img/raspi_logo_16.png) Raspberry PI and utilize it as a IoT server.

[Tutorials](./docs/tutorials/) section has more tutorials.

## Contributing

We welcome and encourage community contributions to documents and examples for other developers. Typo and broken link fixes are appreciated.

--------------

[^1]: [TPCx-IoT Performance Results](https://www.tpc.org/tpcx-iot/results/tpcxiot_perf_results5.asp?version=2)
{: .fs-1}

