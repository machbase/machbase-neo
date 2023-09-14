---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Time
nav_order: 03
permalink: /docs/tutorials/tql/tql-time
---

# TQL Time
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-12 12:00:00 123:456:789'), 10);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-13 12:00:00 123:456:789'), 11);
> ```


TQL은 여러가지 `Time` 조작 함수를 제공한다. 

## Time function

`time` 함수는 입력된 값에 따라 Time 값을 반환한다.

### Now

`now` 인자를 주면 현재 시간을 반환한다.

```js
SQL(`select to_char(time), value from example where time < ?`, time('now'))
CSV()
```

`result`

```
2021-08-12 12:00:00 123:456:789,10
2021-08-13 12:00:00 123:456:789,11
```

### Timestamp

`timestamp` 인자를 주면 해당 timestamp에 맞는 time 값을 반환한다. 

```js
SQL(`select to_char(time), value from example where time = ?`, time(1628737200123456789))
CSV()
```

## Format

`Time` 값이 출력될 때 출력 형식을 지정할 수 있다. 

### None

```js
SQL(`select to_char(time), time from example`)
CSV()
```

`result`

```
2021-08-12 12:00:00 123:456:789,1628737200123456789
2021-08-13 12:00:00 123:456:789,1628823600123456789
```

### Default

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12 03:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-13 03:00:00.123
```

### Numeric

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('NUMERIC'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,08/12 03:00:00AM '21 +0000
2021-08-13 12:00:00 123:456:789,08/13 03:00:00AM '21 +0000
```

### Ansic

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('ANSIC'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,Thu Aug 12 03:00:00 2021
2021-08-13 12:00:00 123:456:789,Fri Aug 13 03:00:00 2021
```

### Unix

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('NUMERIC'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,Thu Aug 12 03:00:00 UTC 2021
2021-08-13 12:00:00 123:456:789,Fri Aug 13 03:00:00 UTC 2021
```

### RFC822

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('RFC822'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,12 Aug 21 03:00 UTC
2021-08-13 12:00:00 123:456:789,13 Aug 21 03:00 UTC
```

### RFC3339

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('RFC3339'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12T03:00:00Z
2021-08-13 12:00:00 123:456:789,2021-08-13T03:00:00Z
```

## Timezone

`tz` 함수를 사용해서 Timezone을 설정할 수 있다.

### Local

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'), tz('local'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12 12:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-13 12:00:00.123
```

### UTC

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'), tz('UTC'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12 03:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-13 03:00:00.123
```

### Seoul

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'), tz('Asia/Seoul'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12 12:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-13 12:00:00.123
```

### EST

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'), tz('EST'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-11 22:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-12 22:00:00.123
```

### Paris

```js
SQL(`select to_char(time), time from example`)
CSV(timeformat('DEFAULT'), tz('Europe/Paris'))
```

`result`

```
2021-08-12 12:00:00 123:456:789,2021-08-12 05:00:00.123
2021-08-13 12:00:00 123:456:789,2021-08-13 05:00:00.123
```
