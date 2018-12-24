package okex

import "time"

type Instrument []struct {
	BaseCurrency   string `json:"base_currency"`
	BaseIncrement  string `json:"base_increment"`
	BaseMinSize    string `json:"base_min_size"`
	InstrumentID   string `json:"instrument_id"`
	MinSize        string `json:"min_size"`
	ProductID      string `json:"product_id"`
	QuoteCurrency  string `json:"quote_currency"`
	QuoteIncrement string `json:"quote_increment"`
	SizeIncrement  string `json:"size_increment"`
	TickSize       string `json:"tick_size"`
}

type Depth struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp time.Time  `json:"timestamp"`
}

type Ticker []struct {
	BestAsk        string    `json:"best_ask"`
	BestBid        string    `json:"best_bid"`
	InstrumentID   string    `json:"instrument_id"`
	ProductID      string    `json:"product_id"`
	Last           string    `json:"last"`
	Ask            string    `json:"ask"`
	Bid            string    `json:"bid"`
	Open24H        string    `json:"open_24h"`
	High24H        string    `json:"high_24h"`
	Low24H         string    `json:"low_24h"`
	BaseVolume24H  string    `json:"base_volume_24h"`
	Timestamp      time.Time `json:"timestamp"`
	QuoteVolume24H string    `json:"quote_volume_24h"`
}

type Trade []struct {
	Time      time.Time `json:"time"`
	Timestamp time.Time `json:"timestamp"`
	TradeID   string    `json:"trade_id"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Side      string    `json:"side"`
}

type Candle []struct {
	Close  string    `json:"close"`
	High   string    `json:"high"`
	Low    string    `json:"low"`
	Open   string    `json:"open"`
	Time   time.Time `json:"time"`
	Volume string    `json:"volume"`
}
