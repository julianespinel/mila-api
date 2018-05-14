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

type MilaClient interface {
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

func InitClient(client *http.Client) MilaClient {
	return Client{httpClient: client}
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
