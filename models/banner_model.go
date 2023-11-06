package models

import (
	"Goblog/global"
	"Goblog/models/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	gorm.Model
	Path      string          `json:"path"`
	Hash      string          `json:"hash"`
	Name      string          `json:"name"`
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片类型 本地 or 七牛
}

// BeforeDelete 删除之前 的钩子函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//本地图片， 删除 还要删除本地存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error("删除本地图片失败", err)
			return err
		}
	}
	return nil
}
