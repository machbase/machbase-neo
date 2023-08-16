---
title: Chart in Python
parent: Tutorial Waves
layout: default
order: 30
nav_order: 230
---

# Chart in Python with Matplotlib

```python
import requests
import json
import datetime
import matplotlib.pyplot as plt
import numpy as np

url = "http://127.0.0.1:5654/db/query"
querystring = {"q":"select * from example order by time limit 200"} 
response = requests.request("GET", url, params=querystring)
data = json.loads(response.text)

sinTs, sinSeries, cosTs, cosSeries = [], [], [], []
for row in data["data"]["rows"]:
    ts = datetime.datetime.fromtimestamp(row[1]/1000000000)
    if row[0] == 'wave.cos':
        cosTs.append(ts)
        cosSeries.append(row[2])
    else:
        sinTs.append(ts)
        sinSeries.append(row[2])

plt.plot(sinTs, sinSeries, label="sin")
plt.plot(cosTs, cosSeries, label="cos")
plt.title("Tutorial Waves")
plt.legend()
plt.show()
```

![](../img/python-chart.jpg)