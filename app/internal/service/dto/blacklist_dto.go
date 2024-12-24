package dto

import "time"

type CreateBlacklistDTO struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	IDCard  string `json:"idCard" binding:"required"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Remark  string `json:"remark"`
}

type UpdateBlacklistDTO struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	IDCard  string `json:"idCard"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Remark  string `json:"remark"`
	Status  int    `json:"status"`
}

type BlacklistDTO struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	IDCard    string    `json:"idCard"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Remark    string    `json:"remark"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CheckBlacklistDTO struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	IDCard string `json:"idCard"`
}
