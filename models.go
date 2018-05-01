package main

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
	Open decimal.Decimal
	High decimal.Decimal
	Low decimal.Decimal
	Close decimal.Decimal
	AdjClose decimal.Decimal
	Volume int64
}