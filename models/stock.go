package models

import (
	"time"
	"github.com/shopspring/decimal"
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Date time.Time
	Symbol string `gorm:"PRIMARY_KEY"`
	Name string
	Currency string
	Open decimal.Decimal `sql:"type:decimal(20,8);"`
	High decimal.Decimal `sql:"type:decimal(20,8);"`
	Low decimal.Decimal `sql:"type:decimal(20,8);"`
	Close decimal.Decimal `sql:"type:decimal(20,8);"`
	AdjClose decimal.Decimal `sql:"type:decimal(20,8);"`
	Volume int64
}