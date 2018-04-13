package controllers

import (
	"github.com/astaxie/beego/context"
	"XiaodiServer/conf"
	"XiaodiServer/models"
	"XiaodiServer/encrypt"
)

func GetOneToken(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseForm()
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
	token := models.GetOneToken(phone)
	ctx.Output.JSON(models.OneTokenResp{StatusCode: conf.SUCCESS, StatusMsg: conf.SUCCESS_MSG, Token: token}, true, false)
}

func AuthOneToken(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseForm()
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
	token := ctx.Request.Form.Get(conf.TOEKN)
	models.AuthOneToken(phone, token)

	newPass := ctx.Request.Form.Get(conf.USER_PASS)
	newPassDecrypt := encrypt.Base64AesDecrypt(newPass)
	models.UpdatePassByOneToken(phone, newPassDecrypt)
	ctx.Output.JSON(models.BaseResp{StatusCode: conf.SUCCESS, StatusMsg: conf.SUCCESS_MSG}, true, false)
}
