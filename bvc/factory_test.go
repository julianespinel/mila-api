package bvc

import (
	"fmt"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/julianespinel/mila-api/models"
	"github.com/shopspring/decimal"
)

func getRandomString() string {
	return randomdata.Letters(5)
}

func getRandomDecimal(initial int, final int) decimal.Decimal {
	randomFloat := randomdata.Decimal(initial, final)
	return decimal.NewFromFloat(randomFloat)
}

func getRandomInt(initial int, final int) int64 {
	return int64(randomdata.Number(initial, final))
}

func GetRandomStock() models.Stock {
	closePrice := getRandomDecimal(2, 6)
	stock := models.Stock{
		Date:     time.Now(),
		Symbol:   getRandomString(),
		Name:     getRandomString(),
		Currency: getRandomString(),
		Open:     getRandomDecimal(0, 3),
		High:     getRandomDecimal(3, 6),
		Low:      getRandomDecimal(0, 2),
		Close:    closePrice,
		AdjClose: closePrice,
		Volume:   getRandomInt(10, 100),
	}
	return stock
}

func GetTestingStocks(size int) []models.Stock {
	stocks := make([]models.Stock, size)
	for i := 0; i < size; i++ {
		stock := GetRandomStock()
		stocks = append(stocks, stock)
	}
	return stocks
}
