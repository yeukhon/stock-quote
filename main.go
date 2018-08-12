package main

import (
  "encoding/json"
  "net/http"
  "os"
  "time"
  "github.com/yeukhon/stock-quote/lib"
)

var API_KEY string = os.Getenv("ALPHA_ADVANTAGE_API_KEY")
var QUOTE_URL string = "https://www.alphavantage.co/query?function=BATCH_QUOTES_US&symbols=MSFT,FB,NFLX,AMZN,IQ&apikey=" + API_KEY

// Modified from the https://stackoverflow.com/questions/17156371
var client = &http.Client{Timeout: time.Second * 5}
func getQuote(url string, target interface{}) error {
    r, err := client.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}

func main() {
  stock_quote := new(lib.StockQuoteResponse)
  r := getQuote(QUOTE_URL, stock_quote)
  if r != nil {
    panic(r)
  }

  // Instance access to specific stock instead of reading from a list,
  // and print as we go.
  stocks := make(map[string]lib.StockQuote)

  lib.PrintHeader()
  for _, v := range stock_quote.BatchQuote {
    stocks[v.Symbol] = v
    lib.PrintToTerminal(v.Symbol, v)
  }
}
