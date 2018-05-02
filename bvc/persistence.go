package bvc

import (
	"github.com/jinzhu/gorm"
	"github.com/julianespinel/mila-api/models"
)

type Persistence struct {
	db *gorm.DB
}

func InitPersistence(db *gorm.DB) Persistence {
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
