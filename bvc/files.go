package bvc

import (
	"fmt"
	"strconv"
	"time"

	"github.com/julianespinel/mila-api/models"
	"github.com/julianespinel/xls"
	"github.com/shopspring/decimal"
)

const cop = "cop"

func getBVCTemporalFileName() string {
	fileName := "bvc-stocks-%v.xls"
	return fmt.Sprintf(fileName, time.Now().Unix())
}

func getStockFromRow(row *xls.Row) (models.Stock, error) {
	stock := models.Stock{}
	for j := row.FirstCol(); j <= row.LastCol(); j++ {
		cell := row.Col(j)
		if j == 1 {
			volume, err := strconv.ParseInt(cell, 10, 64)
			if err != nil {
				return stock, err
			}
			stock.Volume = volume
		}
		if j == 2 {
			stock.Name = cell
		}
		if j == 4 {
			closePrice, err := decimal.NewFromString(cell)
			if err != nil {
				return stock, err
			}
			stock.Close = closePrice
		}
	}
	stock.Currency = cop
	return stock, nil
}

func getStocksFromBVCFile(filePath string) ([]models.Stock, error) {
	stocks := make([]models.Stock, 0)
	xlFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		return stocks, err
	}
	if sheet := xlFile.GetSheet(0); sheet != nil {
		firstRow := 2 // Why? File first row is blank, second row is the table header.
		for i := firstRow; i <= int(sheet.MaxRow); i++ {
			row := sheet.Row(i)
			stock, err := getStockFromRow(row)
			if err != nil {
				break
			}
			stocks = append(stocks, stock)
		}
	}
	return stocks, err
}
