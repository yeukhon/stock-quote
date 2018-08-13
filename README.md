# stock-quote
Get update about my stocks. An adventure to learn Go programming!


## Step 1
Get your free API from Alpha Advantage: https://www.alphavantage.co/support/#api-key (note there is a 5-minute rate limit).

## Step 2
Export your API key.

```
export ALPHA_ADVANTAGE_API_KEY=??
```

## Step 3
Run `main.go` or compile it!


```
John-Wong:stock-quote jwong$ go run main.go
             open         market       diff
MSFT         109.2600     108.6100     -0.65
AMZN         1898.5000    1907.1800    8.68
NFLX         339.5300     343.8100     4.28
FB           180.0000     181.0000     1
IQ           29.3900      28.6300      -0.76
```

![Screenshot!](static/screenshot.png?raw=true "Screenshot")
