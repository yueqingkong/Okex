package main

import (
	"Okex/plat"
	"log"
)

func main() {
	// 现货
	// spot:= plat.NewOKex()
	// 初始化Key
	// spot.InitKeys("123","www","wwwww")

	// log.Print(spot.Instruments())
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
	// future:= plat.NewOKexFuture()
	// log.Print(future.Position())
	// log.Print(future.InstrumenPosition("EOS-USD-190628"))
	// log.Print(future.Account())
	// log.Print(future.SymbolAccount("btc"))
	// log.Print(future.Times("btc"))
	// log.Print(future.SetTimes("btc","BTC-USD-190628","long",10))
	// log.Print(future.Ledger("eos"))
	// log.Print(future.Instruments())
	// log.Print(future.Ticker("EOS-USD-190517"))
}
