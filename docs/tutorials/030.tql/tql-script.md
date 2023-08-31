---
layout: default
grand_parent: Tutorials
parent: TQL
title: TQL Script
nav_order: 06
permalink: /docs/tutorials/tql/tql-script
---

# TQL Map
{:.no_toc}

1. TOC
{:toc}

{: .important }
> For smooth practice, the following query should be run to prepare tables and data.
> ```sql
> CREATE TAG TABLE IF NOT EXISTS EXAMPLE (NAME VARCHAR(20) PRIMARY KEY, TIME DATETIME BASETIME, VALUE DOUBLE SUMMARIZED);
> ```
>

Supporting script language

1. tengo
 [tengo](https://github.com/d5/tengo) is a Golang like script.
 Supports all builtin pakcages(math, text, times, rand, fmt, json, base64, hex) of tengo except "os" excluded for the security reason.
 And added "context" package for exposing the TQL specific features.

*Syntax*: `SCRIPT({ ... script code... })`

## Insert



