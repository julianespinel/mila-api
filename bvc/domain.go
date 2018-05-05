package bvc

import "time"

type Domain struct {
	client      Client
	persistence Persistence
}

func InitDomain(client Client, persistence Persistence) Domain {
	domain := Domain{
		client:      client,
		persistence: persistence,
	}
	return domain
}

func (domain Domain) UpdateDailyStocks() error {
	stocks := domain.client.GetStocksClosingDataByDate(time.Now())
	return domain.persistence.SaveStocks(stocks)
}
