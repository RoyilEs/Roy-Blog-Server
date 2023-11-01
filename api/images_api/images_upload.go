package images_api

import (
	"Goblog/global"
	"Goblog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"fileName"`   //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

// ImageUploadView 上传图片, 返回图片的url
func (ImageApi) ImageUploadView(c *gin.Context) {
	////上传图片
	//fileHeader, err := c.FormFile("image")
	//多张图片
	form, err := c.MultipartForm()
	if err != nil {
		res.ResultFailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.ResultOkWithMsg("不存在的文件", c)
		return
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	//pathList := strings.Split(global.Config.Upload.Path, "/")
	_, err = os.ReadDir(basePath)
	if err != nil {
		//不存在就创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []FileUploadResponse

	//遍历获取
	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小不能超过%dM, 当前大小为%.2fM", global.Config.Upload.Size, size),
			})
			continue
		}
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "上传失败" + err.Error(),
			})
			continue
		}
		//上传成功
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}
	res.ResultOkWithData(resList, c)
}
