---
title: Read waves by shell
parent: Let's make waves
grand_parent: Tutorials
layout: default
order: 30
nav_order: 200
---

# Reading data by shell

## SQL Query

The simplest way to print out data using "SQL query".

```sh
machbase-neo shell "select * from EXAMPLE order by time desc"
```
```
    #  NAME      TIME(UTC)                   VALUE
───────────────────────────────────────────────────────
    1  wave.sin  2023-01-27 23:29:57.000000  0.207914
    2  wave.cos  2023-01-27 23:29:57.000000  -0.978147
    3  wave.sin  2023-01-27 23:29:56.000000  0.587787
    4  wave.cos  2023-01-27 23:29:56.000000  -0.809016
    ...
```

We executed query by `machbase-neo shell` without `sql` sub-command above example.
It properly printed out result of query which is becuase machbase-neo shell takes `sql` sub-command as default as long as there are no other arguments and flags. This means `machbase-neo shell "select..."` is same with `machbase-neo shell sql "select..."`.

So when we use some flags for executing query, sepcifiy `sql` subcommand explicitly like below.

```sh
machbase-neo shell sql \
    --tz America/Los_Angeles \
    "select * from EXAMPLE order by time desc limit 4"
```

```
 #  NAME      TIME(AMERICA/LOS_ANGELES)   VALUE
────────────────────────────────────────────────────
 1  wave.sin  2023-01-28 06:03:59.000000  0.214839
 2  wave.cos  2023-01-28 06:03:59.000000  -0.976649
 3  wave.cos  2023-01-28 06:03:58.000000  -0.804831
 4  wave.sin  2023-01-28 06:03:58.000000  0.593504
```

Machbase treats all time data in UTC as default.
Use `--tz` option to display time in any time-zone other than 'UTC' like above example. 
This flag accepts 'local' and tz database format (eg: 'Europe/Paris').

```sh
machbase-neo shell sql \
    --tz local \
    "select * from EXAMPLE order by time desc limit 4"
```
```
 #  NAME      TIME(UTC)                   VALUE
────────────────────────────────────────────────────
 1  wave.sin  2023-01-28 14:03:59.000000  0.214839
 2  wave.cos  2023-01-28 14:03:59.000000  -0.976649
 3  wave.cos  2023-01-28 14:03:58.000000  -0.804831
 4  wave.sin  2023-01-28 14:03:58.000000  0.593504
 ```

## Table view

It is also possible browsing query result forward/backward with "walk" command like below.

```sh
machbase-neo shell walk "select * from EXAMPLE order by time desc"
```

Then you can scroll up/down with keyboard, press `ESC` to exit table view.

Press `r` to re-execute query to refresh result, it is particularly useful with query was sorted by `order by time desc` to see the latest values when data is continuously being written.

![img](chart02.jpg)
