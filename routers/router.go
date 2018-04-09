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

	rewardRouter := beego.NewNamespace("/reward",
		beego.NSPost("/send", controllers.AssertToken(controllers.SendReward)),
		beego.NSPost("/show", controllers.AssertToken(controllers.ShowReward)),
		beego.NSPost("/show/my/send", controllers.AssertToken(controllers.ShowRewardMySend)),
		beego.NSPost("/show/my/carry", controllers.AssertToken(controllers.ShowRewardMyCarry)),
		beego.NSPost("/show/xiaodian", controllers.AssertToken(controllers.ShowRewardSortXiaodian)),
		beego.NSPost("/show/keyword", controllers.AssertToken(controllers.ShowRewardKeyword)),
		beego.NSPost("/carry", controllers.AssertToken(controllers.CarryReward)),
		beego.NSPost("/delivery", controllers.AssertToken(controllers.DeliveryReward)),
		beego.NSPost("/finish", controllers.AssertToken(controllers.FinishReward)),
		beego.NSPost("/evaluate", controllers.AssertToken(controllers.Evaluate)),
	)

	beego.Get("/picture/:type/:imgid", controllers.GetPicture)
	beego.AddNamespace(userRouter)
	beego.AddNamespace(rewardRouter)
}
