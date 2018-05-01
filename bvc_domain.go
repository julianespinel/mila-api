package main

import "time"

type BVCDomain struct {
	client BVCClient
	persistence BVCPersistence
}

func (domain BVCDomain) UpdateDailyStocks() error {
	stocks := domain.client.GetStocksClosingDataByDate(time.Now())
	return domain.persistence.SaveStocks(stocks)
}
