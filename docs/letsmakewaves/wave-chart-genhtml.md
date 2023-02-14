---
title: Generating HTML chart
parent: Let's make waves
layout: default
order: 30
nav_order: 220
---

# Generating HTML chart

## HTTP Chart API

HTTP API `/db/chart` generates charts.

```python
from urllib import parse
from IPython.display import display, IFrame

req = parse.urlencode([("tags", "example/wave.sin"), ("tags", "example/wave.cos")])

display(IFrame(f"http://127.0.0.1:5654/db/chart?{req}", width=700, height=400))
```

The url of above example is same as below, open it in your web browser.

```
http://127.0.0.1:5654/db/chart?tags=example%2Fwave.sin&tags=example%2Fwave.cos
```


![img](../img/python_http_chart.jpg)

## Command line

Command line tool generates `chart.html` file that contains a chart.

Execute it and open output html file with your web browser.

```sh
machbase-neo shell chart \
    --range 1m \
    --count 1 \
    --output chart.html \
    --format html \
    --html-title "Let's make waves" \
    EXAMPLE/wave.sin EXAMPLE/wave.cos
```

This command generates "chart.html" file that specified in `--output` like below.

![img](../img/chart-html.jpg)


