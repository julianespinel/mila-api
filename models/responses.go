package models

import "time"

type CurrentDayStocksResponse struct {
	Date       time.Time `json:"date"`
	Country    string    `json:"country"`
	StocksData []Stock   `json:"stocksData"`
}
