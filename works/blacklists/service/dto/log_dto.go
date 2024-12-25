package dto

import "time"

type QueryLogDTO struct {
	ID         int       `json:"id"`
	MerchantID int       `json:"merchantId"`
	Phone      string    `json:"phone"`
	IDCard     string    `json:"idCard"`
	Name       string    `json:"name"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"userAgent"`
	Exists     bool      `json:"exists"`
	CreatedAt  time.Time `json:"createdAt"`
}
