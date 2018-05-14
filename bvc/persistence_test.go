package bvc

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julianespinel/mila-api/models"
	"github.com/stretchr/testify/assert"
)

func getDatabaseConnection(t *testing.T) *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/miladb_test.db")
	assert.Nil(t, err)
	return db
}

func dropStocksTable(db *gorm.DB) {
	db.DropTable(&models.Stock{})
}

func Test_BVCPersistence_saveStocks_success(t *testing.T) {
	// Setup
	db := getDatabaseConnection(t)
	size := 5
	stocks := GetTestingStocks(size, models.Colombia)
	assert.Equal(t, size, len(stocks))
	// Exercise
	bvcPersistence := Persistence{db: db}
	err := bvcPersistence.saveStocks(stocks)
	assert.Nil(t, err)
	assert.Equal(t, size, bvcPersistence.countStocks())
	// Tear down
	dropStocksTable(db)
	assert.Equal(t, 0, bvcPersistence.countStocks())
	db.Close()
}
