---
title: User script in TQL
parent: Let's make waves
layout: default
permalink: /docs/letsmakewaves/tql-script/
order: 30
nav_order: 12
---

# User script in TQL
{: .no_toc}

<!--
1. TOC
{: toc }

## SCRIPT()
-->

Open a new *tql* editor on the web ui and copy the code below and run it.

In this example, `linspace(-4,4,100)` generates an array contains 100 elements which are ranged from -4.0 to 4.0 in every `8/100` step. `meshgrid()` takes two array and produce meshed new array. As result of FAKE() in the example produces an array of 10000 elements (100 x 100 meshed) contains array of two float point numbers.
`SCRIPT()` function takes a code block which enclosed by `{` and `}` and run it for each record.
Users can takes the key and value of the records via `context.key()` and `context.value()` then yield transformed data via `context.yield()` or `context.yieldKey()`.

```js
FAKE(meshgrid(linspace(-4,4,100), linspace(-4,4, 100)))
SCRIPT({
  math := import("math")
  // Define a custom function in the script
  calc := func(a, b) {
    return math.sin(math.pow(a, 2) + math.pow(b, 2)) /
           (math.pow(a, 2) + math.pow(b, 2))
  }
  // Receive values of the record from context
  ctx := import("context")
  values := ctx.value()
  x := values[0]
  y := values[1]
  z := calc(x, y)
  // Yield new value
  //  - yieldKey() build and passes new value with new key to the next step.
  //  - yeild() build and passes new value to the next step with the received key from previous step
  ctx.yieldKey(x, y, z)
})
CHART_LINE3D(
  // chart size in HTML syntax
  size('1000px', '600px'),
  // width, height, depth grids in percentage
  gridSize(100,50,100),
  lineWidth(5), visualMap(-0.1, 1),
  // rotation speed in degree per sec.
  autoRotate(20)
)
```

![web-tql-script-wave](/assets/img/web-tql-script-wave.gif)