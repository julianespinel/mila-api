package bvc

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julianespinel/mila-api/files"
	"github.com/julianespinel/mila-api/models"
)

type MilaClient interface {
	GetStocksClosingDataByDate(date time.Time) ([]models.Stock, error)
}

type Client struct {
	httpClient *http.Client
}

const (
	results        = 100
	variableIncome = 1
)

func InitClient(client *http.Client) MilaClient {
	return Client{httpClient: client}
}

func (bvcClient Client) GetStocksClosingDataByDate(date time.Time) ([]models.Stock, error) {
	var stocks []models.Stock
	url := fmt.Sprintf(
		"https://www.bvc.com.co/mercados/DescargaXlsServlet?archivo=acciones&fecha=%s&resultados=%v&tipoMercado=%v",
		date.Format("2006-01-02"),
		results,
		variableIncome,
	)
	log.Println("performing request: ", url)
	res, err := bvcClient.httpClient.Get(url)
	if err != nil {
		return stocks, err
	}
	defer res.Body.Close()
	filePath := getBVCTemporalFileName()
	if err = files.SaveBodyToFile(filePath, res.Body); err != nil {
		return stocks, err
	}
	if stocks, err = getStocksFromBVCFile(filePath); err != nil {
		return stocks, err
	}
	os.Remove(filePath)
	return stocks, err
}
