package dto

import "time"

type LoginLogDTO struct {
	ID        int       `json:"id"`
	Type      int       `json:"type"`
	UserID    int       `json:"userId"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
