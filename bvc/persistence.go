package bvc

import (
	"github.com/jinzhu/gorm"
	"../models"
)

type BVCPersistence struct {
	db *gorm.DB
}

func (persistence BVCPersistence) SaveStocks(stocks []models.Stock) error {
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
