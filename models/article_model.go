package models

import (
	"Goblog/models/ctype"
	"gorm.io/gorm"
)

type ArticleModel struct {
	gorm.Model
	Title      string      `gorm:"size:32" json:"title"`                    //文章标题
	Abstract   string      `gorm:"size:128" json:"abstract"`                //文章介绍
	Content    string      `json:"content"`                                 //文章内容
	Tag        []Tag       `gorm:"many2many:article_tag_models" json:"tag"` //文章标签
	Category   string      `gorm:"size:20" json:"category"`                 //文章分类
	User       User        `gorm:"foreignKey:UserID" json:"-"`              //文章作者
	UserID     uint        `json:"user_id"`                                 //作者ID
	Source     string      `gorm:"size:20" json:"source"`                   //文章来源
	Link       string      `json:"link"`                                    //原文连接
	Word       string      `json:"word"`                                    //字数
	Banner     BannerModel `gorm:"foreignKey:BannerID" json:"banner"`       //封面
	BannerID   uint        `json:"banner_id"`                               //封面ID
	Nickname   string      `gorm:"size:42" json:"nickname"`
	BannerPath string      `json:"banner_path"`
	Tags       ctype.Array `gorm:"type:string;size:64" json:"tags"` //文章标签
}
