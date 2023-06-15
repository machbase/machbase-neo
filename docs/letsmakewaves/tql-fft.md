---
title: FFT in TQL
parent: Let's make waves
layout: default
permalink: /docs/letsmakewaves/tql-fft/
order: 30
nav_order: 11
---

# Fast Fourier Transform in TQL
{: .no_toc}

1. TOC
{: toc }

## Generates sample data

Open a new *tql* editor on the web ui and copy the code below and run it.

In this example, `oscilator()` generates a composite wave of 15Hz 1.0 + 24Hz 1.5.
And `CHART_SCATTER()` has `dataZoom()` option function that provides an slider under the x-Axis.

```js
INPUT( FAKE( 
  oscilator(
    freq(15, 1.0), freq(24, 1.5),
    range('now', '10s', '1ms')) 
  )
)
OUTPUT( CHART_SCATTER( dataZoom('slider', 95, 100)) )
```

![web-fft-tql-fake](/assets/img/web-fft-tql-fake.jpg)

## Store data into database

Store the generated data into the database with the tag name 'signal'.

```js
INPUT( FAKE( oscilator(freq(1.5, 1.0), range('now', '3s', '10ms')) ))
OUTPUT( INSERT( 'time', 'value', table('example'), tag('signal') ) )
```

It will show "10000 rows inserted." message in the "Result" pane.

## Read data from database

The code below is read the stored data from 'example' table.

```js
INPUT( QUERY('value', from('example', 'signal'), between('last-10s', 'last')) )
OUTPUT( CHART_LINE( dataZoom('slider', 95, 100)) )
```

![web-fft-tql-query](/assets/img/web-fft-tql-query.jpg)

## Fast Fourier Transform

Add few data manipulation function between `INPUT()` and `OUTPUT()`.

```js
INPUT( QUERY('value', from('example', 'signal'), between('last-10s', 'last')) )

PUSHKEY('sample')
GROUPBYKEY()
FFT()
FLATTEN()
POPKEY()

OUTPUT(
  CHART_LINE(
        xAxis(0, 'Hz'),
        yAxis(1, 'Amplitude'),
        dataZoom('slider', 0, 10) 
    )
)
```

![web-fft-tql-2d](/assets/img/web-fft-tql-2d.jpg)

## Adding time axis

```js
INPUT( 
  QUERY('value', 
        from('example', 'signal'), 
        between('last-10s', 'last')
  ) 
)

PUSHKEY( roundTime(K, '500ms') )
GROUPBYKEY()
FFT(minHz(0), maxHz(100))

OUTPUT(
  CHART_BAR3D(
        xAxis(0, 'time', 'time'),
        yAxis(1, 'Hz'),
        zAxis(2, 'Amp'),
        size('600px', '600px'), visualMap(0, 1.5), theme('westeros')
  )
)
```

![web-fft-tql-3d](/assets/img/web-fft-tql-3d.jpg)


