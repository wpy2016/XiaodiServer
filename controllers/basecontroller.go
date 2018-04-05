package controllers

import (
	"XiaodiServer/models/resp"
	"fmt"
	"github.com/astaxie/beego/context"
)

func GetBaseErrorResp(code int, msg string) resp.BaseResp {
	resp := resp.BaseResp{}
	resp.StatusCode = code
	resp.StatusMsg = msg
	return resp
}

func CatchErr(ctx *context.Context) {
	{
		if err := recover(); nil != err {
			_, ok := err.(resp.BaseResp)
			if ok {
				ctx.Output.JSON(err, true, false)
				return
			}
			errStr := fmt.Sprintf("%s", err)
			errNew := resp.BaseResp{
				StatusCode: 444,
				StatusMsg:  errStr,
			}
			ctx.Output.JSON(errNew, true, false)
		}
	}
}
