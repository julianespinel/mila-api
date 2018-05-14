package bvc

import (
	"log"
	"strings"
	time "time"

	"github.com/julianespinel/mila-api/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

type API struct {
	domain MilaDomain
}

func InitAPI(domain MilaDomain) API {
	return API{domain: domain}
}

func (api API) AddRoutes(router router.Party) router.Party {
	router.Get("/{country:string}", api.GetCurrentDayStocksByCountry)
	return router
}

func getCurrentDayStocksResponse(country string, stocks []models.Stock) models.CurrentDayStocksResponse {
	return models.CurrentDayStocksResponse{
		Date:       time.Now(),
		Country:    country,
		StocksData: stocks,
	}
}

func (api API) GetCurrentDayStocksByCountry(ctx iris.Context) {
	country := ctx.Params().Get("country")
	if strings.EqualFold(country, models.Colombia) {
		stocks := api.domain.getCurrentDayStocks(country)
		currentDayStocksResponse := getCurrentDayStocksResponse(country, stocks)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(currentDayStocksResponse)
	}
}

func (api API) UpdateDailyStocks(date time.Time) error {
	err := api.domain.updateDailyStocks(date)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
