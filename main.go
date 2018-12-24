package main

import (
	"time"
	"log"
	"OkexApi/okex"
)

func main() {
	log.Print(okex.Candles("BTC-USDT", time.Time{}))
}
