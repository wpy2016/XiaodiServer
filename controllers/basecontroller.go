package controllers

import (
	"XiaodiServer/models"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"XiaodiServer/conf"
)

func GetBaseErrorResp(code int, msg string) models.BaseResp {
	resp := models.BaseResp{}
	resp.StatusCode = code
	resp.StatusMsg = msg
	return resp
}

func CatchErr(ctx *context.Context) {
	{
		if err := recover(); nil != err {
			_, ok := err.(models.BaseResp)
			if ok {
				ctx.Output.JSON(err, true, false)
				return
			}
			errStr := fmt.Sprintf("%s", err)
			errNew := models.BaseResp{
				StatusCode: 444,
				StatusMsg:  errStr,
			}
			ctx.Output.JSON(errNew, true, false)
		}
	}
}

func AssertToken(f beego.FilterFunc) beego.FilterFunc {
	return func (ctx *context.Context) {
		defer CatchErr(ctx)
		ctx.Request.ParseForm()
		userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
		token := ctx.Request.Form.Get(conf.TOEKN)
		models.AssertTokenExist(userId, token)
		if nil != f{
			f(ctx)
		}
	}
}