package bvc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/extrame/xls"
	"github.com/julianespinel/mila-api/models"
	"github.com/shopspring/decimal"
)

type ClientMila interface {
	getStocksClosingDataByDate(date time.Time) []models.Stock
}

type Client struct {
	err        error
	httpClient *http.Client
}

const (
	results        = 100
	variableIncome = 1
	cop            = "cop"
)

/*
 * TODO: improve error handling.
 * See: https://blog.golang.org/errors-are-values
 */

func InitClient(client *http.Client) ClientMila {
	return Client{httpClient: client}
}

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
	stock.Currency = cop
	return stock
}

func getStocksFromFile(filePath string) []models.Stock {
	stocks := make([]models.Stock, 0)
	xlFile, err := xls.Open(filePath, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	if sheet := xlFile.GetSheet(0); sheet != nil {
		firstRow := 2 // Why? File first row is blank, second row is the table header.
		for i := firstRow; i <= int(sheet.MaxRow); i++ {
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

func (bvcClient Client) getStocksClosingDataByDate(date time.Time) []models.Stock {
	url := fmt.Sprintf(
		"https://www.bvc.com.co/mercados/DescargaXlsServlet?archivo=acciones&fecha=%s&resultados=%v&tipoMercado=%v",
		date.Format("2006-01-02"),
		results,
		variableIncome,
	)
	log.Println("performing request: ", url)
	res, err := bvcClient.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	filePath := saveBodyToFile(res.Body)
	stocks := getStocksFromFile(filePath)
	deleteFile(filePath)
	return stocks
}
