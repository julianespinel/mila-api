package core

import (
	"github.com/jinzhu/gorm"
	"github.com/julianespinel/mila-api/models"
)

type MilaPersistence interface {
	countStocks() int
	saveStocks(stocks []models.Stock) error
	getCurrentDayStocks(country string) []models.Stock
}

type Persistence struct {
	db *gorm.DB
}

func InitPersistence(db *gorm.DB) MilaPersistence {
	return Persistence{db: db}
}

func (persistence Persistence) saveStocks(stocks []models.Stock) error {
	persistence.db.CreateTable(&models.Stock{})
	tx := persistence.db.Begin()
	for _, stock := range stocks {
		err := tx.Create(&stock).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (persistence Persistence) countStocks() int {
	count := 0
	persistence.db.Model(&models.Stock{}).Count(&count)
	return count
}

func (persistence Persistence) getCurrentDayStocks(country string) []models.Stock {
	var stocks []models.Stock
	persistence.db.Where(&models.Stock{Country: country}).Find(&stocks)
	return stocks
}
