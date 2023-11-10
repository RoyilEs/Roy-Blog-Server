package models

import (
	"Goblog/models/ctype"
	"gorm.io/gorm"
)

// MenuModel 菜单表
type MenuModel struct {
	gorm.Model
	MenuTitle    string        `json:"menu_title" gorm:"type:varchar(32);comment:菜单名称"`
	MenuTitleEn  string        `json:"menu_title_en" gorm:"type:varchar(32);comment:菜单名称(en)"`
	Slogan       string        `json:"slogan" gorm:"type:varchar(64);comment:菜单标语"`
	Abstract     ctype.Array   `json:"abstract" gorm:"size:64"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `json:"banners" gorm:"many2many:menu_image_models;joinForeignKey:MenuID;joinReferences:BannerID;comment:菜单图标"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `json:"sort" gorm:"size:10"`
}
