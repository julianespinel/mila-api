package models

import (
	"time"
	"github.com/shopspring/decimal"
)

type Stock struct {
	Date time.Time
	Symbol string
	Name string
	Currency string
	Open decimal.Decimal
	High decimal.Decimal
	Low decimal.Decimal
	Close decimal.Decimal
	AdjClose decimal.Decimal
	Volume int64
}