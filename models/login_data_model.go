package models

import "gorm.io/gorm"

type LoginDataModels struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"-"`
	IP       string `gorm:"size:20" json:"ip"`
	Nickname string `gorm:"size:42" json:"nickname"`
	Token    string `gorm:"size:256" json:"token"`
	Device   string `gorm:"size:256" json:"device"`
	Addr     string `gorm:"size:644" json:"addr"`
}
