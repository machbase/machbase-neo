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
    EXAMPLE/wave.sin EXAMPLE/wave.cos
```

In this example, we specified `-` instead of filename for `--output` to prints output on stdout.

```json
{
  "type": "line",
  "data": {
    "labels": [
      "06:33:23","06:33:24","06:33:25","06:33:26","06:33:27","06:33:28","06:33:29",
      "06:33:30","06:33:31","06:33:32","06:33:33","06:33:34","06:33:35","06:33:36"
    ],
    "datasets": [
      {
        "label": "wave.sin",
        "data": [
          -0.587439,-0.865811,-0.994477,-0.951189,-0.743432,-0.407129,-0.00043,
          0.406344,0.742857,0.950923,0.994567,0.866241,0.588135,0.208335
        ],
        "borderWidth": 1
      },
      {
        "label": "wave.cos",
        "data": [
          -0.809268,-0.500371,-0.104955,0.308609,0.668812,0.913371,1,
          0.91372,0.669451,0.309427,-0.1041,-0.499626,-0.808763,-0.978058
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

