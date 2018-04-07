package controllers

import (
	"XiaodiServer/conf"
	"XiaodiServer/models"
	"fmt"
	"github.com/astaxie/beego/context"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func SendReward(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseMultipartForm(1 << 21)
	fmt.Println(ctx.Request.Form)
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
	token := ctx.Request.Form.Get(conf.TOEKN)
	models.AssertTokenExist(userId, token)
	baseUser := models.GetBaseUserById(userId)
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
	xiaodian := ctx.Request.Form.Get(conf.REWARD_XIAODIAN)
	xiaodianInt, err := strconv.ParseInt(xiaodian, 10, 32)
	if nil != err {
		panic(models.BaseResp{conf.ERROR_XIAODIAN_LAYOUT, conf.ERROR_XIAODIAN_LAYOUT_MSG})
	}
	deadlineStr := ctx.Request.Form.Get(conf.REWARD_DEADLINE)
	deadlineTime, err := time.Parse(conf.TIME_FORMAT, deadlineStr)
	if nil != err {
		panic(models.BaseResp{conf.ERROR_TIME_LAYOUT, conf.ERROR_TIME_LAYOUT_MSG})
	}
	originLocation := ctx.Request.Form.Get(conf.REWARD_ORIGIN_LOCATION)
	dstLocation := ctx.Request.Form.Get(conf.REWARD_DST_LOCATION)
	descibe := ctx.Request.Form.Get(conf.REWARD_DESCRIBE)
	thing := getThing(ctx)
	reward := &models.Reward{
		ID:             bson.NewObjectId().Hex(),
		Publisher:      *baseUser,
		State:          conf.REWARD_SEND,
		Phone:          phone,
		Xiaodian:       int(xiaodianInt),
		DeadLine:       deadlineTime,
		OriginLocation: originLocation,
		DstLocation:    dstLocation,
		Describe:       descibe,
		Thing:          thing,
		CreateTime:     time.Now(),
	}
	reward.Save()
	ctx.Output.JSON(models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}, true, false)
}

func ShowReward(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	token := ctx.Request.Form.Get(conf.TOEKN)
	models.AssertTokenExist(userId, token)

	pages := ctx.Request.Form.Get(conf.REWARD_PAGES)
	pagesInt, _ := strconv.Atoi(pages)
	rewards := models.ShowReward(pagesInt)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}

	ctx.Output.JSON(rewardResp, true, false)
}

func CarryReward(ctx *context.Context) {
	defer CatchErr(ctx)
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	token := ctx.Request.Form.Get(conf.TOEKN)
	models.AssertTokenExist(userId, token)

	rewardId := ctx.Request.Form.Get(conf.ID)
	models.CarryReward(rewardId,userId)
	baseResp := models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}
	ctx.Output.JSON(baseResp, true, false)
}

func DeliveryReward(ctx *context.Context) {

}

func FinishReward(ctx *context.Context) {

}

func Evaluate(ctx *context.Context) {

}

func getThing(ctx *context.Context) models.Thing {
	thingType := ctx.Request.Form.Get(conf.THING_TYPE)
	thingTypeInt, err := strconv.ParseInt(thingType, 10, 32)
	if nil != err {
		panic(models.BaseResp{conf.ERROR_THING_TYPE, conf.ERROR_THING_TYPE_MSG})
	}
	var thumbnail string
	file, header, e := ctx.Request.FormFile(conf.THING_THUMBNAIL)
	if nil != e {
		switch thingTypeInt {
		case conf.THING_TYPE_EXPRESS:
			thumbnail = conf.DEFAULT_THING_EXPRESS_IMG
		case conf.THING_TYPE_FOOD:
			thumbnail = conf.DEFAULT_THING_FOOD_IMG
		case conf.THING_TYPE_PAPER:
			thumbnail = conf.DEFAULT_THING_PAPER_IMG
		case conf.THING_TYPE_OTHER:
			thumbnail = conf.DEFAULT_THING_OTHER_IMG
		}
	} else {
		filename, _ := uploadFile(file, header, conf.UPLOAD_IMG_REWARD_FILE_PATH)
		thumbnail = conf.IMG_REWARD_HTTP + filename
	}
	weight := ctx.Request.Form.Get(conf.THING_WEIGHT)
	thing := models.Thing{}
	thing.Thumbnail = thumbnail
	thing.Weight = weight
	thing.ThingType = int(thingTypeInt)
	return thing
}
