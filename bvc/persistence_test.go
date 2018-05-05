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

func TestBVCPersistence_SaveStocks(t *testing.T) {
	// Setup
	db := getDatabaseConnection(t)
	size := 5
	stocks := GetTestingStocks(size)
	// Exercise
	bvcPersistence := Persistence{db: db}
	err := bvcPersistence.SaveStocks(stocks)
	assert.Nil(t, err)
	// Tear down
	dropStocksTable(db)
	db.Close()
}