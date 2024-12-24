package dto

import "time"

type CreateBlacklistDTO struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	IDCard     string `json:"id_card" binding:"required"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	Remark     string `json:"remark"`
	MerchantID int    `json:"merchant_id" binding:"required"`
}

type UpdateBlacklistDTO struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	IDCard  string `json:"id_card"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Remark  string `json:"remark"`
	Status  int    `json:"status"`
}

type BlacklistDTO struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Phone      string    `json:"phone"`
	IDCard     string    `json:"id_card"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	Remark     string    `json:"remark"`
	Status     int       `json:"status"`
	MerchantID int       `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CheckBlacklistDTO struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	IDCard string `json:"id_card"`
}
