---
parent: Tutorials
title: Wave data and monitoring
layout: default
has_children: true
---

# Wave data and monitoring

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

You could delete the table when you've done with it.

```sh
machbase-neo shell "drop table EXAMPLE"
```

## Make waves

We are going to cover serveral different ways how to write data into Machbase-neo through this tutorial.

- [x] Shell script using `machbase-neo shell` command
- [ ] Python client program writing data via HTTP API.
- [x] Go client program writing data via HTTP API.
- [x] Go client program writing data via gRPC API.


## Watch waves

Since we learned various ways to to write data, it's time to know how we can "watch" the data that we collect.

- [x] Graph on terminal
- [X] Table view on terminal
- [x] Generate chart HTML 
- [ ] Using Grapana plugin for Machbase-neo
