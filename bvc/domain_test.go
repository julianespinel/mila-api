package bvc

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initializeBVCDomain(t *testing.T) (Domain, *MockClientInterface,
	*MockPersistenceInterface) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	clientMock := NewMockClientInterface(mockController)
	persistenceMock := NewMockPersistenceInterface(mockController)
	domain := InitDomain(clientMock, persistenceMock)
	return domain, clientMock, persistenceMock
}

func Test_BVCDomain_UpdateDailyStocks_success(t *testing.T) {
	domain, clientMock, persistenceMock := initializeBVCDomain(t)

	size := 5
	stocks := GetTestingStocks(size)
	date := time.Date(2018, time.April, 30, 0, 0, 0, 0, time.UTC)

	clientMock.EXPECT().GetStocksClosingDataByDate(date).Return(stocks)
	persistenceMock.EXPECT().SaveStocks(stocks).Return(nil)

	err := domain.UpdateDailyStocks(date)
	assert.Nil(t, err)
}
