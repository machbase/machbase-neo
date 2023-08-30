---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Basic
nav_order: 00
permalink: /docs/tutorials/tql/tql-basic
---

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
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-14'), 12);
> INSERT INTO EXAMPLE VALUES('TAG0', TO_DATE('2021-08-15'), 13);
> ```
>

## Output

아래의 `Output` 함수들을 사용해서 출력 데이터를 원하는 형식으로 출력할 수 있다.

{: .note }
> `Output` 함수를 확인하기 위해 Source Table로 부터 데이터를 추출하도록 한다. <br/>
> `SQL` 함수를 사용해서 TQL 내부에서 query를 수행할 수 있다.
>
> ```js
> SQL( 'select * from example' )
> ```

### CSV Format

```js
SQL( 'select * from example' )
CSV()
```

### JSON Format

```js
SQL( 'select * from example' )
JSON()
```

### Markdown Format

```js
SQL( 'select * from example' )
MARKDOWN()
```

## TQL API
TQL를 저장해서 HTTP 통신의 Endpoint로 사용할 수 있다.

Save this code as `output.tql`

```js
SQL( 'select * from example' )
JSON()
```

Open it with web browser at [http://127.0.0.1:5654/db/tql/output.tql](http://127.0.0.1:5654/db/tql/output.tql), or use *curl* command on the terminal.

## Input

아래의 `Input` 함수들을 사용해서 데이터를 테이블에 입력할 수 있다.

{: .note }
> `Input` 함수를 확인하기 위해 `Fake` 함수를 사용해서 임의의 데이터를 생성한다.
>
> ```js
> FAKE( oscillator(freq(1.5, 1.0), range('now', '1s', '500ms')) )
> ```

### Insert

`INSERT` function stores incoming records into specified databse table by an `INSERT` statement for each record.

```js
FAKE( oscillator(freq(1.5, 1.0), range('now', '1s', '500ms')) )
INSERT("time", "value", table("example"), tag('temperature'))
```

