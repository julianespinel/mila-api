package bvc

import (
	"encoding/json"
	"testing"
	time "time"

	gomock "github.com/golang/mock/gomock"
	"github.com/julianespinel/mila-api/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/assert"
)

func initializeBVCDomainMock(t *testing.T) *MockDomainMila {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	return NewMockDomainMila(mockController)
}

func newIrisTestApp(api API) *iris.Application {
	app := iris.New()
	milaAPIRoutes := app.Party("/mila/api")
	milaAPIRoutes = api.AddRoutes(milaAPIRoutes)
	return app
}

func Test_BVCAPI_GetCurrentDayStocksByCountry_success(t *testing.T) {
	domainMock := initializeBVCDomainMock(t)
	size := 5
	stocks := GetTestingStocks(size, models.Colombia)
	domainMock.EXPECT().getCurrentDayStocks(models.Colombia).Return(stocks)
	api := InitAPI(domainMock)

	app := newIrisTestApp(api)
	testClient := httptest.New(t, app)
	response := testClient.GET("/mila/api/{country}", models.Colombia).Expect()
	response.Status(httptest.StatusOK)
	// Convert JSON to CurrentDayStocksResponse struct
	jsonString := response.Body().Raw()
	var stocksResponse models.CurrentDayStocksResponse
	json.Unmarshal([]byte(jsonString), &stocksResponse)
	// Check strunct is correct
	assert.NotZero(t, stocksResponse.Date)
	assert.Equal(t, stocksResponse.Country, models.Colombia)
	assert.Len(t, stocksResponse.StocksData, size)
}

func Test_BVCAPI_updateDailyStocks_success(t *testing.T) {
	domainMock := initializeBVCDomainMock(t)
	api := InitAPI(domainMock)

	date := time.Date(2018, time.April, 30, 0, 0, 0, 0, time.UTC)
	domainMock.EXPECT().updateDailyStocks(date).Return(nil)

	err := api.UpdateDailyStocks(date)
	assert.Nil(t, err)
}
