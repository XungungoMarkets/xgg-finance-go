# xgg-finance-go

Fork of [piquette/finance-go](https://github.com/piquette/finance-go) maintained by [XungungoMarkets](https://github.com/XungungoMarkets) with updated Yahoo Finance API support.

## Summary

Go package for accessing current and historical financial markets data via Yahoo Finance.

### Features

Description | Source
--- | ---
Quote(s) | Yahoo finance
Equity quote(s) | Yahoo finance
Index quote(s) | Yahoo finance
Option quote(s) | Yahoo finance
Forex pair quote(s) | Yahoo finance
Cryptocurrency pair quote(s) | Yahoo finance
Futures quote(s) | Yahoo finance
ETF quote(s) | Yahoo finance
Mutual fund quote(s) | Yahoo finance
Historical quotes | Yahoo finance
Options straddles | Yahoo finance

## Installation

```sh
go get github.com/XungungoMarkets/xgg-finance-go
```

## Usage example

### Quote
```go
q, err := quote.Get("AAPL")
if err != nil {
  panic(err)
}
fmt.Println(q)
```

### Historical quotes (OHLCV)
```go
params := &chart.Params{
  Symbol:   "NVDA",
  Interval: datetime.OneHour,
}
iter := chart.Get(params)

for iter.Next() {
  fmt.Println(iter.Bar())
}
if err := iter.Err(); err != nil {
  fmt.Println(err)
}
```

## License

See [LICENSE](LICENSE).
