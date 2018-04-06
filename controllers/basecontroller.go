package controllers

import (
	"XiaodiServer/models"
	"fmt"
	"github.com/astaxie/beego/context"
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
