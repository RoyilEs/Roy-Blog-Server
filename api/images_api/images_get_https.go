package images_api

import (
	"Goblog/global"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (ImageApi) ImageGetHttpsView(c *gin.Context) {
	response, err := http.Get("https://api.vvhan.com/api/acgimg?type=json")
	if err != nil {
		global.Log.Error(err)
		return
	}
	reader := response.Body
	defer func(reader io.ReadCloser) {
		err := reader.Close()
		if err != nil {

		}
	}(reader)
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	c.DataFromReader(200, contentLength, contentType, reader, map[string]string{})
}
