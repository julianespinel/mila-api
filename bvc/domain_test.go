package bvc

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/julianespinel/mila-api/models"
	"github.com/stretchr/testify/assert"
)

func initializeBVCDomain(t *testing.T) (DomainMila, *MockClientMila,
	*MockPersistenceMila) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	clientMock := NewMockClientMila(mockController)
	persistenceMock := NewMockPersistenceMila(mockController)
	domain := InitDomain(clientMock, persistenceMock)
	return domain, clientMock, persistenceMock
}

func Test_BVCDomain_updateDailyStocks_success(t *testing.T) {
	domain, clientMock, persistenceMock := initializeBVCDomain(t)

	size := 5
	stocks := GetTestingStocks(size, models.Colombia)
	date := time.Date(2018, time.April, 30, 0, 0, 0, 0, time.UTC)

	clientMock.EXPECT().getStocksClosingDataByDate(date).Return(stocks)
	persistenceMock.EXPECT().saveStocks(stocks).Return(nil)

	err := domain.updateDailyStocks(date)
	assert.Nil(t, err)
}
