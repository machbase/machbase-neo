---
title: Reading data on terminal
parent: Wave data and monitoring
grand_parent: Tutorials
layout: default
order: 30
nav_order: 210
---

# Reading data on terminal

## Chart on terminal

machabse-neo shell provides simple command line tool for monitoring incoming data.

```sh
machbase-neo shell chart --range 30s EXAMPLE/wave.sin#value EXAMPLE/wave.cos#value
```

![img](chart01.jpg)


## Table view

It is also possible browsing query result forward/backward with "walk" command like below.

```sh
machbase-neo shell walk "select * from EXAMPLE order by time desc"
```

Then you can scroll up/down with keyboard, press `ESC` to exit table view.

![img](chart02.jpg)

{: .note }

> Machbase treats all time data in UTC as default.
> Use `--tz` option with shell command to display time in a time-zone other than 'UTC'. 
> This option accepts 'local' and tz database format (eg: 'Europe/Paris').
> 
> 
> `machbase-neo shell --tz=local walk select...`
>
> `machbase-neo shell --tz=America/Los_Angeles walk select...`

## SQL Query

The simplest way to print out data using "SQL query".

```sh
$ machbase-neo shell "select * from EXAMPLE order by time desc"
    #  NAME      TIME(UTC)                   VALUE
───────────────────────────────────────────────────────
    1  wave.sin  2023-01-27 23:29:57.000000  0.207914
    2  wave.cos  2023-01-27 23:29:57.000000  -0.978147
    3  wave.sin  2023-01-27 23:29:56.000000  0.587787
    4  wave.cos  2023-01-27 23:29:56.000000  -0.809016
    ...
```

