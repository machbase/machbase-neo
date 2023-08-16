---
title: Generating HTML chart
parent: Tutorial Waves
layout: default
order: 30
nav_order: 220
---

{: .warning-title}
> Deprecated
> 
> *HTTP API* `/db/chart` is deprecated. Use *TQL API* instead.

# Generating HTML chart
{: .no_toc }

1. TOC
{:toc}

## HTTP Chart API

HTTP API `/db/chart` generates charts.

```python
from urllib import parse
from IPython.display import display, IFrame
# Can assign multiple tags with the same name 'tags'
req = parse.urlencode([("tags", "example/wave.sin"), ("tags", "example/wave.cos")])
# Jupyter can display embeded <iframe>
display(IFrame(f"http://127.0.0.1:5654/db/chart?{req}", width=700, height=400))
```

Or you can access the url in your web broser, the URL of above example is as below, open it in your web browser.

```
http://127.0.0.1:5654/db/chart?tags=example%2Fwave.sin&tags=example%2Fwave.cos
```


![img](../img/python_http_chart.jpg)

## Command line generates chart

Command line tool generates `chart.html` file that contains a chart.

Execute it and open output html file with your web browser.

```sh
machbase-neo shell chart \
    --range 1m \
    --count 1 \
    --output chart.html \
    --format html \
    --title "Tutorial Waves" \
    EXAMPLE/wave.sin EXAMPLE/wave.cos
```

This command generates "chart.html" file that specified in `--output` like below.

![img](../img/chart-html.jpg)


