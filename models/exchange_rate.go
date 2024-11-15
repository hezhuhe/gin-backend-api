package models

import "time"

type ExchangeRate struct {
	ID           int       `gorm:"primarykey" json:"id"`
	FromCurrency string    `json:"fromcurrency" binding:"required"`
	ToCurrency   string    `json:"tocurrency" bindig:"required"`
	Rate         float64   `json:"rate" binding:"required"`
	Date         time.Time `json:date`
}
