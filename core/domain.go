package core

import (
	"time"

	"github.com/julianespinel/mila-api/bvc"
	"github.com/julianespinel/mila-api/models"
)

type MilaDomain interface {
	updateDailyStocks(date time.Time) error
	getCurrentDayStocks(country string) []models.Stock
}

type Domain struct {
	bvcClient   bvc.MilaClient
	persistence MilaPersistence
}

func InitDomain(client bvc.MilaClient, persistence MilaPersistence) MilaDomain {
	domain := Domain{
		bvcClient:   client,
		persistence: persistence,
	}
	return domain
}

func (domain Domain) updateDailyStocks(date time.Time) error {
	stocks, err := domain.bvcClient.GetStocksClosingDataByDate(date)
	if err != nil {
		return err
	}
	return domain.persistence.saveStocks(stocks)
}

func (domain Domain) getCurrentDayStocks(country string) []models.Stock {
	return domain.persistence.getCurrentDayStocks(country)
}
