package okex

import (
	"net/http"
	"log"
	"time"
	"encoding/json"
	"fmt"
)

type Candle []struct {
	Close  string    `json:"close"`
	High   string    `json:"high"`
	Low    string    `json:"low"`
	Open   string    `json:"open"`
	Time   time.Time `json:"time"`
	Volume string    `json:"volume"`
}

var okApi = "https://www.okex.com"

func get(url string, inter interface{}) {
	client := &http.Client{}
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = client.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(inter)
	if err != nil {
		log.Fatal(err)
	}
}

func Candles(category string, st time.Time) Candle {
	var start string
	var end string
	if st.IsZero() {
		oldStart, err := time.Parse("2006-01-02 15:04:05", "2017-10-11 00:00:00")
		if err != nil {
			log.Print(err)
		}
		start = oldStart.UTC().Format(time.RFC3339)

		oldEnd, err := time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")
		if err != nil {
			log.Print(err)
		}
		end = oldEnd.UTC().Format(time.RFC3339)
	} else {
		start = st.UTC().Format(time.RFC3339)
		//获取200条日线
		end = st.Add(time.Duration(24) * time.Hour * 200).UTC().Format(time.RFC3339)
	}

	var api = fmt.Sprintf("/api/spot/v3/instruments/%s/candles?granularity=86400&start=%s&end=%s", category, start, end)
	var url = okApi + api
	log.Print(url)

	var cand Candle
	get(url, &cand)

	return cand
}
