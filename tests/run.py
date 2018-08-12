from bottle import route, run

@route('/')
def returnarray():
    return {"username": "john", "logins": [{"date": "06/18"}, {"date": "06/19"}]}

@route("/stock")
def stock():
    return {
    "Meta Data": {
        "1. Information": "US Stock Quotes",
        "2. Notes": "IEX Real-Time",
        "3. Time Zone": "US/Eastern"
    },
    "Stock Batch Quotes": [
        {
            "1. symbol": "MSFT",
            "2. open": "109.4400",
            "3. high": "109.6900",
            "4. low": "108.3800",
            "5. price": "109.0000",
            "6. volume": "18119337",
            "7. timestamp": "2018-08-10 16:00:00 EST",
            "8. currency": "USD"
        },
        {
            "1. symbol": "FB",
            "2. open": "182.0900",
            "3. high": "182.1000",
            "4. low": "179.4200",
            "5. price": "180.2600",
            "6. volume": "21440247",
            "7. timestamp": "2018-08-10 16:00:00 EST",
            "8. currency": "USD"
        },
        {
            "1. symbol": "AAPL",
            "2. open": "207.3100",
            "3. high": "209.1000",
            "4. low": "206.6700",
            "5. price": "207.5300",
            "6. volume": "24845250",
            "7. timestamp": "2018-08-10 16:00:00 EST",
            "8. currency": "USD"
        }
    ]
}
run(host='localhost', port=8080, debug=True)
