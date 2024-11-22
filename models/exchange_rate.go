package models

import "time"

type ExchangeRate struct {
	ID           int       `gorm:"primarykey" json:"id"`
	FromCurrency string    `json:"from_currency" binding:"required"`
	ToCurrency   string    `json:"to_currency" bindig:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:date`
}
