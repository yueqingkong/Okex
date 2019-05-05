package main

import (
	"Okex/plat"
	"log"
)

func main() {
	// 现货
	spot:= plat.NewOKex()
	log.Print(spot.Instruments())
	// log.Print(spot.Depths("eth"))
	// log.Print(spot.Tickers())
	// log.Print(spot.Depths("btc"))
	// log.Print(spot.Trades("btc"))
	// log.Print(spot.Accounts())
	// log.Print(spot.Ledger("btc"))
	// log.Print(spot.Order("eos","buy",1,0))
	// log.Print(spot.CancelOrder("eos","2771597172357120"))
	// log.Print(spot.OrderList("eos",0))
	// log.Print(spot.OrderPending())
	// log.Print(spot.OrderInfo("eos","2771624839041024"))
	// log.Print(spot.Deals("eos","2771624839041024"))


	// 合约
	//log.Print(api.FutureCandle("","btc", time.Time{}))
}
