package bvc

import "time"

type Domain struct {
	client      ClientInterface
	persistence PersistenceInterface
}

func InitDomain(client ClientInterface, persistence PersistenceInterface) Domain {
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
