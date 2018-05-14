package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type Stock struct {
	gorm.Model `json:"-"`      // Ignore gorm fields: ID, CreatedAt, UpdatedAt.
	Date       time.Time       `json:"date"`
	Country    string          `json:"country"`
	Symbol     string          `gorm:"PRIMARY_KEY" json:"symbol"`
	Name       string          `json:"name"`
	Currency   string          `json:"currency"`
	Open       decimal.Decimal `sql:"type:decimal(20,8);" json:"open"`
	High       decimal.Decimal `sql:"type:decimal(20,8);" json:"high"`
	Low        decimal.Decimal `sql:"type:decimal(20,8);" json:"low"`
	Close      decimal.Decimal `sql:"type:decimal(20,8);" json:"close"`
	AdjClose   decimal.Decimal `sql:"type:decimal(20,8);" json:"adjClose"`
	Volume     int64           `json:"volume"`
}
