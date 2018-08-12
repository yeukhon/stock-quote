package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
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
  fmt.Println("hi")
  stock_quote := new(lib.StockQuoteResponse)
  getQuote(QUOTE_URL, stock_quote)
  fmt.Println(stock_quote.Metadata)
}
