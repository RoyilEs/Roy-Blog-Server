package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID          uint   `json:"id"`
	MenuTitle   string `json:"menu_title"`
	MenuTitleEn string `json:"menu_title_en"`
	Path        string `json:"path"`
}

func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Model(models.MenuModel{}).Select("id", "menu_title", "menu_title_en", "path").Scan(&menuNameList)
	res.ResultOkWithData(menuNameList, c)
}
