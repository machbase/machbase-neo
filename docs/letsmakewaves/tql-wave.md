---
title: Waves in TQL
parent: Let's make waves
layout: default
permalink: /docs/letsmakewaves/tql-wave/
order: 30
nav_order: 10
---

# Let's make waves in TQL
{: .no_toc}

1. TOC
{: toc }


Open a new *tql* editor on the web ui.

![web-tql](/assets/img/web-tql.jpg)

## Generates data

Copy the code below and run it.

```js
INPUT( FAKE( oscilator(freq(1.5, 1.0), range('now', '3s', '10ms')) ))
OUTPUT( CHART_LINE() )
```

![web-hello-tql](/assets/img/web-hello-tql.jpg)

## Store data into database

Replace `OUTPUT( CHART_LINE() )` with `OUTPUT( INSERT() )` as the code below storing the generated data into the database.

```js
INPUT( FAKE( oscilator(freq(1.5, 1.0), range('now', '3s', '10ms')) ))
OUTPUT( INSERT( 'time', 'value', table('example'), tag('wave.sin') ) )
```

![web-hello-tql-insert](/assets/img/web-hello-tql-insert.jpg)

## Read data from database

### Using `SQL()`

```js
INPUT( SQL(`select time, value from example where name = 'wave.sin'`) )
OUTPUT( CHART_SCATTER() )
```

![web-hello-tql-sql](/assets/img/web-hello-tql-sql.jpg)

### Using `QUERY()`

```js
INPUT( QUERY('value', from('example', 'wave.sin'), between('last-3s', 'last')) )
OUTPUT( CHART_BAR() )
```

![web-hello-tql-query](/assets/img/web-hello-tql-query.jpg)

The `INPUT( QUERY(...) )` function above generates a SQL statement internally like below. It just saves developer's time making a SQL query with proper time conditions.

```sql
SELECT
    time, value 
FROM EXAMPLE
WHERE
    name = 'wave.sin'
AND time BETWEEN
        (SELECT MAX_TIME-3000000000 FROM V$EXAMPLE_STAT WHERE name = 'wave.sin')
    AND (SELECT MAX_TIME FROM V$EXAMPLE_STAT WHERE name = 'wave.sin')
LIMIT 0, 1000000
```

## TQL as RESTful API - HTTP GET

Any *tql* script that saved as a file in '.tql' extension can be invoked via HTTP GET request. Save the example code above as `hello.tql` then open it with your web brower at [http://127.0.0.1:5654/db/tql/hello.tql](http://127.0.0.1:5654/db/tql/hello.tql) or use `curl -o - http://127.0.0.1:5654/db/tql/hello.tql` on a terminal.

### JSON()

```js
INPUT( QUERY('value', from('example', 'wave.sin'), between('last-3s', 'last')) )
OUTPUT( JSON() )
```

### JSON() with transpose()

```js
INPUT( QUERY('value', from('example', 'wave.sin'), between('last-3s', 'last')) )
OUTPUT( JSON( transpose(true) ) )
```

### Query param

```js
INPUT( QUERY('value', from('example', 'wave.sin'), between('last-3s', 'last')) )
OUTPUT( JSON( transpose( $trans ?? false) ) )
```

- [http://127.0.0.1:5654/db/tql/hello.tql?trans=true](http://127.0.0.1:5654/db/tql/hello.tql?trans=true)
- [http://127.0.0.1:5654/db/tql/hello.tql?trans=false](http://127.0.0.1:5654/db/tql/hello.tql?trans=false)

### CSV()

```js
INPUT( QUERY('value', from('example', 'wave.sin'), between('last-3s', 'last')) )
OUTPUT( CSV() )
```

### CSV() with pandas

Load CSV data from *tql* to pandas in Python.

```python
from urllib import parse
import pandas as pd

path = "http://127.0.0.1:5654/db/tql/hello.tql"
df = pd.read_csv(path, header=None)
df
```

![web-hello-tql-csv-pandas](/assets/img/web-hello-tql-csv-pandas.jpg)

```python
from urllib import parse
import pandas as pd

path = "http://127.0.0.1:5654/db/tql/hello.tql"
df = pd.read_csv(path, names=['time', 'amplitude'], header=None)
df['time']= pd.to_datetime(df['time'])
df.plot(x='time', y='amplitude')
```

![web-hello-tql-csv-pandas-plot](/assets/img/web-hello-tql-csv-pandas-plot.jpg)

## TQL as RESTful API - HTTP POST

A *tql* that starts with `INPUT(CSV(CTX.Body...))` then ends with `OUTPUT(APPEND(...))` works as an API that ingests data.

Make a new script and save it as `append.tql`.

```js
INPUT(
    CSV(
        CTX.Body,
        col(0, stringType(), 'name'),
        col(1, datetimeType('s'), 'time'),
        col(2, doubleType(), 'value')
    )
)
OUTPUT( APPEND(table('example')) )
```

{:.note}
> The 'APPEND' works only when fields of input records exactly match with columns of the table in order and types.

And make test data in CSV `data.csv`

```
barn,1677646800,0.03135
dew_point,1677646800,24.4
dishwasher,1677646800,3.33e-05
fridge,1677646800,0.12415
furnace,1677646800,0.0207
garage_door,1677646800,0.0130833
gen,1677646800,0.00348333
home_office,1677646800,0.442633
house_overall,1677646800,0.932833
humidity,1677646800,0.62
```

Send the csv file to the `append.tql` via HTTP POST.

```sh
curl -o - -v -X POST --data-binary @data.csv http://127.0.0.1:5654/db/tql/append.tql
```

```sh
$ curl -o - -v -X POST --data-binary @data.csv http://127.0.0.1:5654/db/tql/append.tql
> POST /db/tql/append2.tql HTTP/1.1
> Host: 127.0.0.1:5654
> User-Agent: curl/7.88.1
> Accept: */*
> Content-Length: 283
> Content-Type: application/x-www-form-urlencoded
>
< HTTP/1.1 200 OK
< Content-Type: application/octet-stream
< Date: Fri, 16 Jun 2023 06:15:11 GMT
< Content-Length: 36
<
append 10 rows (success 10, fail 0).
```
