package main

import (
	"OKex/plat"
	"log"
)

func main() {
	log.Print(plat.Instruments())
	//log.Print(plat.Depths("BTC-USDT"))
	//log.Print(plat.Tickers())
	//log.Print(plat.Depths("BTC-USDT"))
	//log.Print(plat.Trades("BTC-USDT"))
	//log.Print(plat.Candles("BTC-USDT", time.Time{}))
}
