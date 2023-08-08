---
layout: page
title: Roadmap
permalink: /features/
nav_order: 10
parent: Welcome to machbase
---

# Roadmap

## 2023 Q1

- [x] HTTP Server
- [x] MQTT Server

## 2023 Q2

- [x] MQTT Subscriber

## 2023 2H

- [ ] Kafka Consumer
- [ ] NATS Subscriber

```mermaid
flowchart LR
    subgraph machbase-neo
        machbase[("machbase engine")]
        machbase <--Read/Write--o http["http listener ✓Done"]
        machbase <--Write--o mqtt["mqtt listener ✓Done"]
        machbase <--Write--o mqttsub["mqtt subscriber ✓Done"]
        machbase <--Write--o kafkasub["kafka subscriber ☐Todo"]
        machbase <--Write--o natssub["nats subscriber ☐Todo"]
    end
    subgraph external systems
        http x--request--o httpclient["http client"]
        mqtt x--publish--o mqttclient["mqtt client"]
        mqttsub --subscribe--> mqttbroker["mqtt broker"]
        kafkasub --consume--> kafkabroker["kafka server"]
        natssub --subscribe--> natsbroker["nats server"]
    end
```

