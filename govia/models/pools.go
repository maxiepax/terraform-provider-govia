package models

import "time"

type Pool struct {
	ID               int       `json:"id"`
	Name             string    `json:"name" gorm:"type:varchar(255);not null" binding:"required" `
	NetAddress       string    `json:"net_address" gorm:"type:varchar(15);not null"`
	StartAddress     string    `json:"start_address" gorm:"type:varchar(15);not null" binding:"required" `
	EndAddress       string    `json:"end_address" gorm:"type:varchar(15);not null" binding:"required" `
	Netmask          int       `json:"netmask" gorm:"type:integer;not null" binding:"required" `
	LeaseTime        int       `json:"lease_time" gorm:"type:bigint" binding:"required" `
	Gateway          string    `json:"gateway" gorm:"type:varchar(15)" binding:"required" `
	OnlyServeReimage bool      `json:"only_serve_reimage" gorm:"type:boolean"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
