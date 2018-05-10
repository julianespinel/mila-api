package bvc

import (
	"github.com/jinzhu/gorm"
	"github.com/julianespinel/mila-api/models"
)

type PersistenceInterface interface {
	CountStocks() int
	SaveStocks(stocks []models.Stock) error
}

type Persistence struct {
	db *gorm.DB
}

func InitPersistence(db *gorm.DB) PersistenceInterface {
	return Persistence{db: db}
}

func (persistence Persistence) SaveStocks(stocks []models.Stock) error {
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

func (persistence Persistence) CountStocks() int {
	count := 0
	persistence.db.Model(&models.Stock{}).Count(&count)
	return count
}
