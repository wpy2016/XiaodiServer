package controllers

import (
	"XiaodiServer/conf"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetPicture(ctx *context.Context) {
	filename := ctx.Input.Param(":imgid")
	currentPath, err := getCurrentPath()
	filePath := currentPath + conf.UPLOAD_IMG_HEAD_FILE_PATH + string(os.PathSeparator) + filename
	fmt.Println(filePath)
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.Output.SetStatus(404)
		return
	}
	file, err := os.Open(filePath)
	if nil != err {
		ctx.Output.SetStatus(404)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		ctx.Output.SetStatus(404)
		return
	}

	ctx.ResponseWriter.Header().Add("Content-Type", "image/jpeg")
	suffix := GetSuffix(filename)
	if ".png" == suffix {
		png.Encode(ctx.ResponseWriter, img)
		return
	}
	if ".jpg" == suffix || ".jpeg" == suffix {
		jpeg.Encode(ctx.ResponseWriter, img, nil)
	}
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
