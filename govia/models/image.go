package models

import "time"

type Image struct {
	ID          int       `json:"id"`
	ISOImage    string    `json:"iso_image" gorm:"type:varchar(255)"`
	Hash        string    `json:"hash" gorm:"-"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
