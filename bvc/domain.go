package bvc

import (
	"time"

	"github.com/julianespinel/mila-api/models"
)

type DomainMila interface {
	updateDailyStocks(date time.Time) error
	getCurrentDayStocks(country string) []models.Stock
}

type Domain struct {
	client      ClientInterface
	persistence PersistenceMila
}

func InitDomain(client ClientInterface, persistence PersistenceMila) DomainMila {
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
