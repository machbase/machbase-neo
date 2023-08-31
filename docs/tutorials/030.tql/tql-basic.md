<!-- TODO ---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Basic
nav_order: 00
permalink: /docs/tutorials/tql/tql-basic
--- -->

# TQL Basic
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-12'), 10);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-13'), 11);
> ```
>

## Output

아래의 `Output` 함수들을 사용해서 출력 데이터를 원하는 형식으로 출력할 수 있다.

{: .note }
> `Output` 함수를 확인하기 위해 Source Table로 부터 데이터를 추출하도록 한다. <br/>
> `SQL` 함수를 사용해서 TQL 내부에서 query를 수행할 수 있다.
>
> ```js
> SQL( `select * from example` )
> ```

### CSV Format

```js
SQL( `select * from example` )
CSV()
```

`result`

```
TAG0,1628694000000000000,10
TAG0,1628780400000000000,11
```

### JSON Format

```js
SQL( `select * from example` )
JSON()
```

`result`

```json
{
  "data": {
    "columns": [
      "NAME",
      "TIME",
      "VALUE"
    ],
    "types": [
      "string",
      "datetime",
      "double"
    ],
    "rows": [
      [
        "TAG0",
        1628694000000000000,
        10
      ],
      [
        "TAG0",
        1628780400000000000,
        11
      ]
    ]
  },
  "success": true,
  "reason": "success",
  "elapse": "1.571444ms"
}
```

### Markdown Format

```js
SQL( `select * from example` )
MARKDOWN()
```

`result`

```
|NAME|TIME|VALUE|
|:-----|:-----|:-----|
|TAG0|1628694000000000000|10.000000|
|TAG0|1628780400000000000|11.000000|
```

## Input

아래의 `Input` 함수들을 사용해서 데이터를 테이블에 입력할 수 있다.

### Insert

`INSERT` function stores incoming records into specified databse table by an `INSERT` statement for each record.

```js
SQL( `select * from example` )
INSERT("name", "time", "value", table("example"))
```

`result`

```
"2 rows inserted."
```

### Append

`APPEND` function stores incoming records into specified databse table via the `APPEND` method of machbase-neo.

- `table()` table(string) specify destination table

```js
SQL( `select * from example` )
APPEND(table('example'))
```

`result`

```
"append 2 rows (success 2, fail 0)."
```
