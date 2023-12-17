package user_api

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/res"
	"Goblog/utils"
	"Goblog/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"fileName"`   //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

// UserUpdateAvatarView 用户头像变更
func (UserApi) UserUpdateAvatarView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	fileHeader, err := c.FormFile("avatar")
	if err != nil {
		res.ResultFailWithMsg(err.Error(), c)
		return
	}

	//判断路径是否存在
	basePathAvatar := global.Config.Upload.PathAvatar
	baseSizeAvatar := global.Config.Upload.SizeAvatar
	_, err = os.ReadDir(basePathAvatar)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(basePathAvatar, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 获取文件名
	fileName := fileHeader.Filename
	nameList := strings.Split(fileName, ".")
	// 获取后缀 全部小写
	suffix := strings.ToLower(nameList[len(nameList)-1])

	if !utils.InList(suffix, global.WhiteImageList) {
		res.ResultFail(FileUploadResponse{
			FileName:  fileName,
			IsSuccess: false,
			Msg:       "不支持的文件类型",
		}, "", c)
		return
	}
	filePath := path.Join(basePathAvatar, fileName)
	// 判断大小
	size := float64(fileHeader.Size) / float64(1024*1024)
	if size >= float64(baseSizeAvatar) {
		res.ResultFail(FileUploadResponse{
			FileName:  fileName,
			IsSuccess: false,
			Msg:       fmt.Sprintf("文件大小不能超过%dM, 当前大小为%.2fM", baseSizeAvatar, size),
		}, "", c)
		return
	}
	fileObj, err := fileHeader.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, _ := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 数据库查询这个图片是否存在
	var (
		avatarModel models.AvatarModel
		userModel   models.User
	)
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		res.ResultFailWithMsg("用户不存在", c)
		return
	}
	err = global.DB.Take(&avatarModel, "hash = ?", imageHash).Error
	if err == nil {
		err = global.DB.Model(&userModel).Update("avatar", avatarModel.Path).Error
		if err != nil {
			res.ResultFailWithMsg("头像修改失败", c)
			return
		}
		res.ResultOkWithMsg("头像修改成功", c)
		return
	}
	// 存储
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		global.Log.Error(err)
		res.ResultFail(FileUploadResponse{
			FileName:  fileName,
			IsSuccess: false,
			Msg:       "图片上传失败",
		}, "", c)
		return
	}
	err = global.DB.Model(&userModel).Update("avatar", avatarModel.Path).Error
	if err != nil {
		res.ResultFailWithMsg("头像修改失败", c)
		return
	}
	res.ResultOkWithMsg("头像修改成功", c)

	//图片入库
	global.DB.Create(&models.AvatarModel{
		Path: "/" + filePath,
		Hash: imageHash,
		Name: fileName,
	})
}
