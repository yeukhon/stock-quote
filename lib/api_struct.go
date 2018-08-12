/*
I don't mean being a dickhead because the free API offering is awesome. I am
grateful for that,but the JSON response is just horrifying. The CSV response
is, however, much easier to parse. Please do something about the JSON response!
*/

package lib

type StockQuoteResponse struct {
  Metadata ResponseMetadata `json:"Meta Data"`
  BatchQuote BatchQuoteResponse `json:"Stock Batch Quotes"`
}

type ResponseMetadata struct {
  Information string `json:"1. Information"`
  Notes       string `json:"2. Notes"`
  TimeZone    string `json:"3. Time Zone"`
}

type BatchQuoteResponse []StockQuote
//  Quotes []StockQuote `json:"Stock Batch Quotes"`
//}

type StockQuote struct {
  Symbol    string `json:"1. symbol"`
  Open      string `json:"2. open"`
  High      string `json:"3. high"`
  Low       string `json:"4. low"`
  Price     string `json:"5. price"`
  Volume    string `json:"6. volume"`
  Last_TS   string `json:"7. timestamp"`
  Currency  string `json:"8. currency"`
}

type Local struct {
  Username string `json:"username"`
  Logins   []Login `json:"logins"`
}

type Login struct {
  Date string `json:"date"`
}
