---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Map
nav_order: 04
permalink: /docs/tutorials/tql/tql-map
---

# TQL Map
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-12'), 10);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-13'), 11);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-14'), 12);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-15'), 13);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-16'), 14);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-17'), 15);
> ```
>

TQL supports serveral `Map` functions.

## TAKE

Takes first `n` records and stop the stream.

```js
SQL(`select * from example`)
TAKE(2)
CSV()
```

`result`

```
TAG0,1628694000000000000,10
TAG0,1628780400000000000,11
```

## DROP

Ignore first `n` records, it simply drops the n records.

```js
SQL(`select * from example`)
DROP(2)
CSV()
```

`result`

```
TAG0,1628866800000000000,12
TAG0,1628953200000000000,13
TAG0,1629039600000000000,14
TAG0,1629126000000000000,15
```

## PUSHKEY

Apply new key on each record. The orignal key is push into value tuple.

```js
SQL(`select * from example`)
PUSHKEY('neo')
CSV()
```

`result`

```
neo,TAG0,1628694000000000000,10
neo,TAG0,1628780400000000000,11
neo,TAG0,1628866800000000000,12
neo,TAG0,1628953200000000000,13
neo,TAG0,1629039600000000000,14
neo,TAG0,1629126000000000000,15
```

## POPKEY

Drop current key of the record, then promote *idx*th element of *tuple* as a new key.

```js
SQL(`select * from example`)
POPKEY()
CSV()
```

`result`

```
1628694000000000000,10
1628780400000000000,11
1628866800000000000,12
1628953200000000000,13
1629039600000000000,14
1629126000000000000,15
```

## GROUPBYKEY

Takes multiple continuous records that have same key, then produces a new record which have value array contains all individual values.

For example, if an original records was `{key:k, value:[v1, v2]}`, `{key:k, value:{v3, v4}}`...`{key:k, value:{vx, vy}}`, it produces the new record as `{key:k, value:[[v1,v2],[v3,v4],...,[vx,vy]]}`.

```js
SQL(`select value from example limit 3`)
PUSHKEY('neo')
CSV()
```

`result`

```
neo,10
neo,11
neo,12
```

`apply GROUPBYKEY function`

```js
SQL(`select value from example limit 3`)
PUSHKEY('neo')
GROUPBYKEY()
CSV()
```

`result`

```
neo,[]interface {},[]interface {},[]interface {}
```

## FLATTEN

It works the oposite way of GROUPBYKEY(). Take a record whose value is multi-dimension tuple, produces multiple records for each elements of the tuple reducing the dimension.

For example, if an original record was `{key:k, value:[[v1,v2],[v3,v4],...,[vx,vy]]}`, it produces the new multiple records as `{key:k, value:[v1, v2]}`, `{key:k, value:{v3, v4}}`...`{key:k, value:{vx, vy}}`.

```js
SQL(`select value from example limit 3`)
PUSHKEY('neo')
GROUPBYKEY()
FLATTEN()
CSV()
```

`result`

```
neo,10
neo,11
neo,12
```

## FILTER

Apply the condition statement on the incoming record, then it pass the record only if the condition is true.

For example, if an original record was `{key: k1, value[v1, v2]}` and apply `FILTER(count(V) > 2)`, it simply drop the record. If the codition was `FILTER(count(V) >= 2)`, it pass the record to the next function.

```js
SQL(`select name, value from example`)
FILTER(value(0) < 12)
CSV()
```

`result`

```
TAG0,10
TAG1,11
```

## FFT(Fast Fourier Transform)

It assumes value of the incoming record is an array of *time,amplitude* tuples, then applies *Fast Fourier Transform* on the array and replaces the value with an array of *frequency,amplitude* tuples. The key remains same.

For example, if the incoming record was `{key: k, value[ [t1,a1],[t2,a2],...[tn,an] ]}`, it transforms the value to `{key:k, value[ [F1,A1], [F2,A2],...[Fm,Am] ]}`.

`insert test data`

```js
FAKE( oscillator(
    freq(15, 1.0), freq(24, 1.5),
    range('now', '10s', '1ms')) 
)
PUSHKEY('signal')
APPEND( table('example') )
```

`FFT function example`

```js
SQL(`select time, value from example where name = 'signal'`)

MAPKEY('sample')
GROUPBYKEY()
FFT()
CHART_LINE(
      xAxis(0, 'Hz'),
      yAxis(1, 'Amplitude'),
      dataZoom('slider', 0, 10) 
)
```

![web-fft-tql-2d](/assets/img/web-fft-tql-2d.png)
