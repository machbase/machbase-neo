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

1. TOC
{: toc }

## SCRIPT()

Open a new *tql* editor on the web ui and copy the code below and run it.

In this example, `linspace(-4,4,100)` generates an array contains 100 elements which are from -4.0 to 4.0 in every `8/100` step. `meshgrid()` takes two array and produce meshed new array. As result of FAKE() in the example produces an array of 10000 elements (100 x 100 meshed) contains array of two float point numbers.
`SCRIPT()` function takes a code block which enclosed by `{` and `}` and run it for each record.
Users can takes the key and value of the records via `context.key()` and `context.value()` then yield transformed data via `context.yield()` or `context.yieldKey()`.

```js
INPUT( FAKE(meshgrid(linspace(-4,4,100), linspace(-4,4, 100))) )
SCRIPT({
  math := import("math")
  // Define custom functions
  calc := func(x, y) {
    return math.sin(math.pow(x, 2) + math.pow(y, 2)) /
           (math.pow(x, 2) + math.pow(y, 2))
  }
  ctx := import("context")
  // Receive values of the record from context
  x := ctx.value()[0]
  y := ctx.value()[1]
  z := calc(x, y)
  // Yield new values by calculation
  //  - yieldKey() passes values with new key to the next step.
  //  - yeild() passes values with current key to the next step.
  ctx.yieldKey(x, y, z)
})
OUTPUT(
  CHART_LINE3D(
    size('1000px', '600px'), gridSize(100,50,100),
    lineWidth(5), visualMap(-0.1, 1), autoRotate(20)
  )
)
```

![web-tql-script-wave](/assets/img/web-tql-script-wave.gif)