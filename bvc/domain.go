package bvc

import (
	"time"

	"github.com/julianespinel/mila-api/models"
)

type MilaDomain interface {
	updateDailyStocks(date time.Time) error
	getCurrentDayStocks(country string) []models.Stock
}

type Domain struct {
	client      MilaClient
	persistence MilaPersistence
}

func InitDomain(client MilaClient, persistence MilaPersistence) MilaDomain {
	domain := Domain{
		client:      client,
		persistence: persistence,
	}
	return domain
}

func (domain Domain) updateDailyStocks(date time.Time) error {
	stocks := domain.client.getStocksClosingDataByDate(date)
	return domain.persistence.saveStocks(stocks)
}

func (domain Domain) getCurrentDayStocks(country string) []models.Stock {
	return domain.persistence.getCurrentDayStocks(country)
}
