package dto

import "time"

type CreateMerchantDTO struct {
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	Remark        string `json:"remark"`
	IPWhitelist   string `json:"ip_whitelist"`
}

type UpdateMerchantDTO struct {
	ID            int    `json:"id" binding:"required"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	Remark        string `json:"remark"`
	IPWhitelist   string `json:"ip_whitelist"`
}

type MerchantDTO struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	ContactPerson   string    `json:"contact_person"`
	ContactPhone    string    `json:"contact_phone"`
	Remark          string    `json:"remark"`
	Status          int       `json:"status"`
	IPWhitelist     string    `json:"ip_whitelist"`
	APIKey          string    `json:"api_key"`
	APISecret       string    `json:"api_secret"`
	APIToken        string    `json:"api_token"`
	TokenExpireTime time.Time `json:"token_expire_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
