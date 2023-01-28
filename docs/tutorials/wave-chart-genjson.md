---
title: Generating JSON for Chart.js
parent: Wave data and monitoring
grand_parent: Tutorials
layout: default
order: 30
nav_order: 230
---

# Generating JSON form Chart.js

Command below generates JSON data that is compatible for [Chart.js](https://www.chartjs.org/docs/latest/){:target="_blank" rel="noopener"}

```sh
machbase-neo shell chart \
    --range 15s \
    --count 1 \
    --output - \
    --format json \
    EXAMPLE/wave.sin
```

In this example, we specified `-` instead of filename for `--output` to prints output on stdout.

```json
{
  "type": "line",
  "data": {
    "labels": [
      "01:04:14","01:04:15","01:04:16","01:04:17","01:04:18","01:04:19","01:04:20",
      "01:04:21","01:04:22","01:04:23","01:04:24","01:04:25","01:04:26","01:04:27"
    ],
    "datasets": [
      {
        "label": "wave.sin",
        "data": [
          -0.587748,-0.866002,-0.994517,-0.951071,-0.743176,-0.40678,-4.8e-05,
          0.406693,0.743112,0.951041,0.994527,0.86605,0.587826,0.207961
        ],
        "borderWidth": 1
      }
    ]
  },
  "options": {
    "scales": {
      "y": {
        "beginAtZero": false
      }
    }
  }
}
```

