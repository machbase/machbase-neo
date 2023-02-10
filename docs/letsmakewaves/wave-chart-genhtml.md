---
title: Generating HTML chart
parent: Let's make waves
layout: default
order: 30
nav_order: 220
---

# Generating HTML chart

Command below generates `chart.html` file that contains a chart.

Execute it and open output html file with your web browser.

```sh
machbase-neo shell chart \
    --range 1m \
    --count 1 \
    --output chart.html \
    --format html \
    --html-title "Let's make waves" \
    EXAMPLE/wave.sin EXAMPLE/wave.cos
```

This command generates "chart.html" file that specified in `--output` like below.

![img](../img/chart-html.jpg)

