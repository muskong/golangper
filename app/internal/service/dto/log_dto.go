package dto

import "time"

type LoginLogDTO struct {
	ID        int       `json:"id"`
	Type      int       `json:"type"` // 1:商户 2:管理员
	UserID    int       `json:"user_id"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	Status    int       `json:"status"` // 1:成功 2:失败
	CreatedAt time.Time `json:"created_at"`
}

type QueryLogDTO struct {
	ID         int       `json:"id"`
	MerchantID int       `json:"merchant_id"`
	Phone      string    `json:"phone"`
	IDCard     string    `json:"id_card"`
	Name       string    `json:"name"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
	Exists     bool      `json:"exists"`
	CreatedAt  time.Time `json:"created_at"`
}
