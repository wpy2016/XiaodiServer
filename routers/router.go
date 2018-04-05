// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"XiaodiServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
	userRouter := beego.NewNamespace("/user",
		beego.NSPost("/register", controllers.Register),
		beego.NSPost("/login", controllers.Login),
		beego.NSGet("/exist/nickname/:nickname", controllers.IsNickNameExist),
		beego.NSGet("/exist/phone/:phone", controllers.IsPhoneExist),
	)

	beego.Get("/picture/:imgid", controllers.GetPicture)
	beego.AddNamespace(userRouter)
}
