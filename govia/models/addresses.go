package models

import "time"

type Address struct {
	ID           int       `json:"id"`
	IP           string    `json:"ip" gorm:"type:varchar(15);not null;index:uniqIp,unique"`
	Mac          string    `json:"mac" gorm:"type:varchar(17);not null"`
	Hostname     string    `json:"hostname" gorm:"type:varchar(255)"`
	Domain       string    `json:"domain" gorm:"type:varchar(255)"`
	Reimage      bool      `json:"reimage" gorm:"type:bool;index:uniqIp,unique"`
	PoolID       int       `json:"pool_id" gorm:"type:BIGINT" swaggertype:"integer"`
	GroupID      int       `json:"group_id" gorm:"type:BIGINT" swaggertype:"integer"`
	Progress     int       `json:"progress" gorm:"type:INT"`
	Progresstext string    `json:"progresstext" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
