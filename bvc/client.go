package bvc

import (
	"fmt"
	"time"
	"io"
	"os"
	"log"
	"github.com/extrame/xls"
	"strconv"
	"github.com/shopspring/decimal"
	"net/http"
	"crypto/tls"

	"../models"
)

type BVCClient struct {
	err error
}

const RESULTS = 100
const VARIABLE_INCOME = 1
const COP = "COP"

/*
 * TODO: improve error handling.
 * See: https://blog.golang.org/errors-are-values
 */

func getTemporalFileName() string {
	fileName := "bvc-stocks-%v.xls"
	return fmt.Sprintf(fileName, time.Now().Unix())
}

func saveBodyToFile(body io.ReadCloser) string {
	filePath := getTemporalFileName()
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, body)
	if err != nil {
		log.Fatal(err)
	}
	return filePath
}

func getStockFromRow(row *xls.Row) models.Stock {
	stock := models.Stock{}
	for j := row.FirstCol(); j <= row.LastCol(); j++ {
		cell := row.Col(j)
		log.Println(cell)
		if j == 1 {
			volume, err := strconv.ParseInt(cell, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			stock.Volume = volume
		}
		if j == 2 {
			stock.Name = cell
		}
		if j == 4 {
			closePrice, err := decimal.NewFromString(cell)
			if err != nil {
				log.Fatal(err)
			}
			stock.Close = closePrice
		}
	}
	stock.Currency = COP
	return stock
}

func getStocksFromFile(filePath string) []models.Stock {
	stocks := make([]models.Stock, 0)
	xlFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	if sheet := xlFile.GetSheet(0); sheet != nil {
		for i := 2; i <= int(sheet.MaxRow); i++ {
			row := sheet.Row(i)
			stock := getStockFromRow(row)
			stocks = append(stocks, stock)
		}
	}
	return stocks
}

func deleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func (bvcClient BVCClient) GetStocksClosingDataByDate(date time.Time) []models.Stock {
	dateStr := date.Format("2006-01-02")
	url := fmt.Sprintf(
		"https://www.bvc.com.co/mercados/DescargaXlsServlet?archivo=acciones&fecha=%s&resultados=%v&tipoMercado=%v",
		/*dateStr,*/ dateStr,
		RESULTS,
		VARIABLE_INCOME,
	)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	filePath := saveBodyToFile(res.Body)
	stocks := getStocksFromFile(filePath)
	deleteFile(filePath)
	return stocks
}
