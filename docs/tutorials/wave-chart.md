---
title: Let's make waves
parent: Tutorials
layout: default
has_children: true
---

# Let's make waves

Through this tutorial, We are going to cover serveral different ways how to write data into Machbase-neo and read from it.


## Run machbase-neo server

Start machbase-neo server.

```sh
machbase-neo serve
```

## Create example table

Create `EXAMPLE` table for this course if it doesn't exist.

```sh
machbase-neo shell "create tag table EXAMPLE (name varchar(100) primary key, time datetime basetime, value double)"
```

You can delete the table first when you want to create fresh one.

```sh
machbase-neo shell "drop table EXAMPLE"
```

## Make waves - write data

- [x] Shell script using `machbase-neo shell` command
- [ ] Python client program writing data via HTTP API.
- [x] Go client program writing data via HTTP API.
- [x] Go client program writing data via gRPC API.


## Watch waves - read data

- [x] Graph on terminal
- [X] Table view on terminal
- [x] Generate chart HTML 
- [ ] Using Grapana plugin for Machbase-neo
