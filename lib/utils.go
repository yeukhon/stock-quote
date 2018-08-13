package lib

import (
  "fmt"
  "log"
  "math"
  "strings"
  "strconv"
)

// Adopted from https://stackoverflow.com/questions/39544571
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

// Real stupid string to float64 conversion with error handling
func StrToNum(num_str string) float64 {
  value, err := strconv.ParseFloat(num_str, 64)
  if err != nil {
    log.Fatal("Unable to convert %s to float64", num_str)
    panic(err)
  }
  return value
}

// All colorized functions
func Green(num float64) {
  fmt.Printf("%-13s\n", "\x1b[32;1m" + fmt.Sprint(num) + "\x1b[0m")
}

func Red(num float64) {
  fmt.Printf("%-13s\n", "\x1b[31;1m" + fmt.Sprint(num) + "\x1b[0m")
}

func Normal(num float64) {
  fmt.Printf("%-13s\n", fmt.Sprint(num)[:6])
}

// Printer functions
func PrintHeader() {
  fmt.Printf("%-13s", strings.Repeat(" ", 13))
  fmt.Printf("%-13s", "open")
  fmt.Printf("%-13s", "market")
  fmt.Printf("%-13s\n", "diff")
}

func PrintToTerminal(symbol string, stock StockQuote) {
  fmt.Printf("%-13s", symbol)
  fmt.Printf("%-13s", stock.Open)
  fmt.Printf("%-13s", stock.Price)

  // Round up the difference
  _diff := StrToNum(stock.Price) - StrToNum(stock.Open)
  diff := Round(_diff, 0.005)

  // Colorized difference
  if _diff == 0 {
    Normal(diff)
  } else if _diff > 0 {
    Green(diff)
  } else {
    Red(diff)
  }
}
