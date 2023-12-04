package flag

import (
	"Goblog/global"
	"Goblog/models"
	"fmt"
)

func Makemigrations() {
	var err error
	//err = global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.BannerModel{})
	//if err != nil {
	//	global.Log.Error("[ error ] 生成连表失败", err)
	//	return
	//}

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},
			&models.Tag{},
			&models.BannerModel{},
			&models.ArticleModel{},
			&models.MenuImageModel{},
			&models.MenuModel{},
			&models.LoginDataModels{},
			&models.AnosuAll{},
			&models.AvatarModel{},
		)
	if err != nil {
		global.Log.Error("[error] 生成数据库表结构失败", err)
		return
	}
	fmt.Println("Makemigrations Roy [ ok ]")
}
