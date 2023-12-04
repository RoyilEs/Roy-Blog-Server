package user_api

import (
	"github.com/gin-gonic/gin"
)

// UserAvatar TODO 上传头像
type UserAvatar struct {
	UserID uint `json:"user_id" binding:"required" msg:"用户ID错误"`
}

type FileUploadResponse struct {
	FileName  string `json:"fileName"`   //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

const AvatarPath = "/uploads/avatar"

// UserUpdateAvatarView 用户头像变更
func (UserApi) UserUpdateAvatarView(c *gin.Context) {
	//file, err := c.FormFile("image")
	//if err != nil {
	//	res.ResultFailWithMsg(err.Error(), c)
	//	return
	//}
	////判断路径是否存在
	//basePath := AvatarPath
	//baseSize := global.Config.Upload.Size
}
