package models

import (
	"Goblog/models/ctype"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	Nickname      string         `json:"nickname" gorm:"type:varchar(20);not null"`        //昵称
	Username      string         `json:"username" gorm:"type:varchar(20);not null;unique"` //用户名
	Email         string         `json:"email" gorm:"type:varchar(255);"`                  //邮箱
	Phone         string         `json:"phone" gorm:"type:varchar(20);"`                   //手机号
	Password      string         `json:"password" gorm:"type:varchar(255);not null"`       //密码
	Permission    ctype.Role     `json:"permission" gorm:"type:int(10);not null"`          //权限
	Avatar        string         `json:"avatar" gorm:"type:varchar(255);"`                 //头像
	Token         string         `json:"token" gorm:"size:64"`                             //其他平台的唯一地址
	ArticleModels []ArticleModel `json:"-" gorm:"foreignKey:UserID"`                       //发布的文章列表
}
