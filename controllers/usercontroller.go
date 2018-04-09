package controllers

import (
	"XiaodiServer/conf"
	"XiaodiServer/encrypt"
	"XiaodiServer/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gopkg.in/mgo.v2/bson"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type UserController struct {
	beego.Controller
}

func Register(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseMultipartForm(1 << 21)
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
	pass := ctx.Request.Form.Get(conf.USER_PASS)
	nickName := ctx.Request.Form.Get(conf.USER_NICKNAME)

	isPhoneExist, err := models.IsPhoneExist(phone)
	if nil != err || isPhoneExist {
		panic(GetBaseErrorResp(conf.ERROR_PHONE_IS_EXIST, conf.ERROR_PHONE_IS_EXIST_MSG))
		return
	}

	isNickName, err := models.IsNickNameExist(nickName)
	if nil != err || isNickName {
		panic(GetBaseErrorResp(conf.ERROR_NICKNAME_EXIST, conf.ERROR_NICKNAME_EXIST_MSG))
		return
	}

	file, fHead, _ := ctx.Request.FormFile("img")
	imgPath, _ := uploadFile(file, fHead, conf.UPLOAD_IMG_HEAD_FILE_PATH)
	httpImgPath := conf.IMG_HEAD_HTTP + imgPath
	decryptPass := encrypt.Base64AesDecrypt(pass)
	user := models.RegisterDefaultUser(phone, decryptPass, nickName, httpImgPath)
	user.Save()

	token := &models.UserToken{}
	token.UserId = user.ID
	token.Token = user.Token
	token.Save()

	user.Pass = pass
	userResp := &models.UserResp{}
	userResp.StatusCode = conf.SUCCESS
	userResp.StatusMsg = conf.SUCCESS_MSG
	userResp.User = *user
	ctx.Output.JSON(userResp, true, false)
}

func Login(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseForm()
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
	pass := ctx.Request.Form.Get(conf.USER_PASS)
	decryptPass := encrypt.Base64AesDecrypt(pass)
	user := models.Login(phone, decryptPass)

	user.Pass = pass
	userResp := &models.UserResp{}
	userResp.StatusCode = conf.SUCCESS
	userResp.StatusMsg = conf.SUCCESS_MSG
	userResp.User = *user
	ctx.Output.JSON(userResp, true, false)
}

func IsPhoneExist(ctx *context.Context) {
	defer CatchErr(ctx)
	phone := ctx.Input.Param(":phone")
	isUsed, err := models.IsPhoneExist(phone)
	if nil != err {
		ctx.Output.JSON(GetBaseErrorResp(444, err.Error()), true, false)
		return
	}
	if isUsed {
		ctx.Output.JSON(GetBaseErrorResp(conf.ERROR_PHONE_IS_EXIST, conf.ERROR_PHONE_IS_EXIST_MSG), true, false)
		return
	}
	ctx.Output.JSON(GetBaseErrorResp(conf.SUCCESS, conf.SUCCESS_MSG), true, false)
}

func IsNickNameExist(ctx *context.Context) {
	defer CatchErr(ctx)
	nickName := ctx.Input.Param(":nickname")
	isUsed, err := models.IsNickNameExist(nickName)
	if nil != err {
		ctx.Output.JSON(GetBaseErrorResp(444, err.Error()), true, false)
		return
	}
	if isUsed {
		ctx.Output.JSON(GetBaseErrorResp(conf.ERROR_NICKNAME_EXIST, conf.ERROR_NICKNAME_EXIST_MSG), true, false)
		return
	}
	ctx.Output.JSON(GetBaseErrorResp(conf.SUCCESS, conf.SUCCESS_MSG), true, false)
}

func uploadFile(file multipart.File, fHead *multipart.FileHeader, filepath string) (string, error) {
	defer file.Close()
	suffix := GetSuffix(fHead.Filename)
	isValid := isValidSuffix(suffix)
	if !isValid {
		return "", errors.New(string(conf.ERROR_IMG_KIND_TYPE) + conf.ERROR_IMG_KIND_TYPE_MSG)
	}
	filename := bson.NewObjectId().Hex() + suffix

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath, 0777)
		if nil != err {
			return "", err
		}
	}

	saveFilePath := filepath + string(os.PathSeparator) + filename
	f, err := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if nil != err {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if nil != err {
		return "", err
	}
	return filename, nil
}

func GetSuffix(filename string) string {
	splite := strings.Split(filename, ".")
	lenght := len(splite)
	return "." + splite[lenght-1]
}

func isValidSuffix(suffix string) bool {
	if suffix == ".png" || suffix == ".jpg" || suffix == ".jpeg" {
		return true
	}
	return false
}
