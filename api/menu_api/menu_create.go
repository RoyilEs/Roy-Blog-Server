package menu_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/ctype"
	"Goblog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle    string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`
	MenuTitleEn  string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"`
	Slogan       string      `json:"slogan"`
	Abstract     ctype.Array `json:"abstract"`
	AbstractTime int         `json:"abstract_time"`                         //切换时间 单位：s
	BannerTime   int         `json:"banner_time"`                           //切换时间 单位：s
	Sort         int         `json:"sort" binding:"required" msg:"请输入菜单序号"` //菜单排序序号
	//ImageSortList []ImageSort `json:"image_sort_list"`                       //具体图片的顺序
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	//添加
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.ResultFailWithError(err, &cr, c)
		return
	}
	//重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "menu_title = ? or menu_title_en = ?", cr.MenuTitle, cr.MenuTitleEn).RowsAffected
	if count > 0 {
		res.ResultFailWithMsg("重复的菜单信息", c)
		return
	}
	//创建Menu数据入库
	menuModel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		res.ResultFailWithMsg("菜单添加失败", c)
		return
	}
	//批量入库
	//if len(cr.ImageSortList) == 0 {
	//	res.ResultOkWithMsg("菜单添加成功", c)
	//	return
	//}
	//var menuImageList []models.MenuImageModel
	//for _, sort := range cr.ImageSortList {
	//	//TODO 判断imageID 是否真正有这张图片
	//	menuImageList = append(menuImageList, models.MenuImageModel{
	//		MenuID:   menuModel.ID,
	//		BannerID: sort.ImageID,
	//		Sort:     sort.Sort,
	//	})
	//}
	////给第三张表入库
	//err = global.DB.Create(&models.MenuImageModel{}).Error
	//if err != nil {
	//	res.ResultFailWithMsg("菜单图片关联失败", c)
	//	return
	//}
	res.ResultOkWithMsg("菜单添加成功", c)
}
