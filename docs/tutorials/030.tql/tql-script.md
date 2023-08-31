---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Script
nav_order: 06
permalink: /docs/tutorials/tql/tql-script
---

# TQL Script
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-12'), 10);
> INSERT INTO EXAMPLE VALUES('TAG1', TO_DATE('2021-08-13'), 11);
> ```
>

Supporting script language

1. tengo
 [tengo](https://github.com/d5/tengo) is a Golang like script.
 Supports all builtin pakcages(math, text, times, rand, fmt, json, base64, hex) of tengo except "os" excluded for the security reason.
 And added "context" package for exposing the TQL specific features.

*Syntax*: `SCRIPT({ ... script code... })`

## Context

Returns context object of the script runtime.

### yieldKey

인자로 들어온 데이터를 Out stream으로 전달한다.

#### Output CSV

```js
SCRIPT({
    ctx := import("context")
    ctx.yieldKey(0, 1, 2, 3)
    ctx.yieldKey(1, 2, 3, 4)
})
CSV()
```

`result`

```
0,1,2,3
1,2,3,4
```

#### Table Append

```js
SCRIPT({
    ctx := import("context")
    ctx.yieldKey("tag0", 100, 10)
    ctx.yieldKey("tag0", 111, 11)
})
APPEND(table('example'))
```

`result`

```
append 2 rows (success 2, fail 0).
```

### key

Returns the key of the current record.

```js
SQL(`select * from example`)
SCRIPT({
    ctx := import("context")
    ctx.yieldKey(ctx.key(), 0, 1, 2, 3)
})
CSV()
```

`result`

```
TAG0,0,1,2,3
TAG1,0,1,2,3
```

### value

Returns the whole value of the current records in array. If the index is given, it returns the element of the values.

For example, If the current value is `[0, true, "hello", "world"]`

- `value()` returns the whole value array `[0, true, "hello", "world"]`
- `value(0)` returns the first element of the value `0`
- `value(3)` returns the last element of the value `"world"`

```js
SQL(`select * from example`)
SCRIPT({
    ctx := import("context")
    ctx.yieldKey(ctx.value(1), 0, 1, 2, 3)
})
CSV()
```

`result`

```
10,0,1,2,3
11,0,1,2,3
```
