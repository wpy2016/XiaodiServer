package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type UserController struct {
	beego.Controller
}

func Register(ctx *context.Context) {

}

func Login(ctx *context.Context) {

}

func IsPhoneExist(ctx *context.Context) {
	ctx.WriteString("hello world" + ctx.Input.Param(":phone"))
}

func IsNickNameExist(ctx *context.Context) {
	ctx.WriteString("hello world" + ctx.Input.Param(":nickname"))
}
