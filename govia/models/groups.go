package models

import "time"

type Group struct {
	ID          int             `json:"id"`
	PoolID      int             `json:"pool_id"`
	Name        string          `json:"name"`
	Password    string          `json:"password"`
	DNS         string          `json:"dns"`
	Ntp         string          `json:"ntp"`
	ImageID     int             `json:"image_id"`
	Syslog      string          `json:"syslog"`
	Vlan        string          `json:"vlan"`
	CallbackURL string          `json:"callbackurl"`
	BootDisk    string          `json:"bootdisk" gorm:"type:varchar(255)"`
	Options     map[string]bool `json:"options,omitempty"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
