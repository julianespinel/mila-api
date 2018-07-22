package core

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
		if len(stocks) == 0 {
			api.UpdateDailyStocks()
		}
		currentDayStocksResponse := getCurrentDayStocksResponse(country, stocks)
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(currentDayStocksResponse)
	}
}

func getLatestBusinessDay() time.Time {
	now := time.Now()
	weekday := now.Weekday().String()
	switch weekday {
	case "Saturday":
		return now.AddDate(0, 0, -1)
	case "Sunday":
		// return now.AddDate(0, 0, -2)
		return now.AddDate(0, 0, -3) // To get thursday, delete this line
	default:
		return now
	}
}

func (api API) UpdateDailyStocks() error {
	err := api.domain.updateDailyStocks(getLatestBusinessDay())
	if err != nil {
		log.Print(err)
	}
	return err
}
