package global

import (
	"Goblog/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config         *config.Config
	DB             *gorm.DB
	Log            *logrus.Logger
	MySqlLog       logger.Interface
	WhiteImageList = []string{
		"jgp",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
		"bmp",
	} //图片白名单
)
