package currency

import "time"

type Currency struct {
	Symbol      string    `json:"symbol"`
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
	Open        string    `json:"open"`
	Volume      string    `json:"volume"`
	VolumeQuote string    `json:"volume_quote"`
	Timestamp   time.Time `json:"timestamp"`
}

type Currencies []Currency
