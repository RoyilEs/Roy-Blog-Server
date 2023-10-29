package models

import "gorm.io/gorm"

type BannerModel struct {
	gorm.Model
	Path string `json:"path"`
	Hash string `json:"hash"`
	Name string `json:"name"`
}
