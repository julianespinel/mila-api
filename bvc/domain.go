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

func (domain Domain) UpdateDailyStocks(date time.Time) error {
	stocks := domain.client.GetStocksClosingDataByDate(date)
	return domain.persistence.SaveStocks(stocks)
}
