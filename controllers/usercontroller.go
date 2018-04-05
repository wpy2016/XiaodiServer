package controllers

import (
	"XiaodiServer/conf"
	"XiaodiServer/encrypt"
	"XiaodiServer/models"
	"XiaodiServer/models/resp"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"gopkg.in/mgo.v2/bson"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

type UserController struct {
	beego.Controller
}

func Register(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseMultipartForm(32 << 20)
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
	imgPath, _ := uploadFile(file, fHead)
	httpImgPath := conf.IMG_HEAD_HTTP + imgPath
	decryptPass := encrypt.Base64AesDecrypt(pass)
	user := models.RegisterDefaultUser(phone, decryptPass, nickName, httpImgPath)
	user.Save()

	user.Pass = ""
	userResp := &resp.UserResp{}
	userResp.StatusCode = conf.SUCCESS
	userResp.StatusMsg = conf.SUCCESS_MSG
	userResp.User = *user
	ctx.Output.JSON(userResp, true, false)
}

func Login(ctx *context.Context) {

}

func IsPhoneExist(ctx *context.Context) {

}

func IsNickNameExist(ctx *context.Context) {

}

func validateSummaryForm(form url.Values) (map[string]string, error) {
	kv := make(map[string]string)

	return kv, nil
}

func uploadFile(file multipart.File, fHead *multipart.FileHeader) (string, error) {
	defer file.Close()
	suffix := GetSuffix(fHead.Filename)
	isValid := isValidSuffix(suffix)
	if !isValid {
		return "", errors.New(string(conf.ERROR_IMG_KIND_TYPE) + conf.ERROR_IMG_KIND_TYPE_MSG)
	}
	filename := bson.NewObjectId().Hex() + suffix

	_, err := os.Stat(conf.UPLOAD_IMG_HEAD_FILE_PATH)
	if os.IsNotExist(err) {
		err = os.MkdirAll(conf.UPLOAD_IMG_HEAD_FILE_PATH, 0777)
		if nil != err {
			return "", err
		}
	}

	saveFilePath := conf.UPLOAD_IMG_HEAD_FILE_PATH + string(os.PathSeparator) + filename
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
