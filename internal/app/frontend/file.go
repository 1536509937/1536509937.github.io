package frontend

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/text/gstr"
	"1536509937/ku-bbs/internal/service"
	"1536509937/ku-bbs/pkg/config"
	"1536509937/ku-bbs/pkg/utils/encrypt"
)

var File = cFile{}

type cFile struct{}

func (c *cFile) MDUploadSubmit(ctx *gin.Context) {
	s := service.Context(ctx)

	file, err := ctx.FormFile("editormd-image-file")
	if err != nil {
		s.MDFileJson(0, err.Error(), "")
		return
	}

	if file.Size > 1024*1024*config.Conf.Upload.TopicFileSize {
		s.MDFileJson(0, "仅支持小于 2M 大小的图片", "")
		return
	}

	arr := strings.Split(file.Filename, ".")
	ext := arr[len(arr)-1]

	if !gstr.InArray(config.Conf.Upload.ImageExt, ext) {
		s.MDFileJson(0, "file format not supported", "")
		return
	}

	path := fmt.Sprintf("%s/topic", config.Conf.Upload.Path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
		os.Chmod(path, os.ModePerm)
	}

	name := encrypt.Md5(time.Now().String()+file.Filename) + "." + ext

	if err := ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", path, name)); err != nil {
		s.MDFileJson(0, err.Error(), "")
	} else {
		s.MDFileJson(1, "ok", fmt.Sprintf("/assets/upload/topic/%s", name))
	}
}
