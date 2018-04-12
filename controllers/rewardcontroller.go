package controllers

import (
	"XiaodiServer/conf"
	"XiaodiServer/models"
	"github.com/astaxie/beego/context"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

func SendReward(ctx *context.Context) {
	ctx.Request.ParseMultipartForm(1 << 21)
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
	user := models.GetUserById(userId)
	xiaodian := ctx.Request.Form.Get(conf.REWARD_XIAODIAN)
	xiaodianInt, err := strconv.ParseInt(xiaodian, 10, 32)

	//检查笑点是否充足
	if user.GoldMoney+user.SilverMoney < int(xiaodianInt) {
		panic(models.BaseResp{conf.ERROR_USER_XIAODIAN_SHORT, conf.ERROR_USER_XIAODIAN_SHORT_MSG})
	}

	baseUser := models.GetBaseUserById(userId)
	phone := ctx.Request.Form.Get(conf.USER_PHONE)
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
		DeadLine:       deadlineStr,
		DeadLineTime:   deadlineTime,
		OriginLocation: originLocation,
		DstLocation:    dstLocation,
		Describe:       descibe,
		Thing:          thing,
		CreateTime:     time.Now(),
	}
	reward.Save()
	ctx.Output.JSON(models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}, true, false)
}

func UpdateReward(ctx *context.Context) {
	ctx.Request.ParseMultipartForm(1 << 21)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
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

	newReward := &models.Reward{
		ID:             rewardId,
		Publisher:      *baseUser,
		State:          conf.REWARD_SEND,
		Phone:          phone,
		Xiaodian:       int(xiaodianInt),
		DeadLine:       deadlineStr,
		DeadLineTime:   deadlineTime,
		OriginLocation: originLocation,
		DstLocation:    dstLocation,
		Describe:       descibe,
		Thing:          thing,
		CreateTime:     time.Now(),
	}
	models.UpdateReward(rewardId, userId, *newReward)
	ctx.Output.JSON(models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}, true, false)
}

func ShowRewardMySend(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
	rewards := models.ShowRewardMySend(userId)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

/**
userid 发布的
receiveId 领取的
 */
func ShowRewardOurNotFinish(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
	receiveId := ctx.Request.Form.Get(conf.RECEIVE_ID)
	rewards := models.ShowRewardOurNotFinish(userId, receiveId)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func ShowRewardMyCarry(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.REWARD_RECEIVER_USER_ID)
	rewards := models.ShowRewardMyCarry(userId)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func ShowRewardMyFinish(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	rewards := models.ShowRewardMyFinish(userId)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func ShowReward(ctx *context.Context) {
	ctx.Request.ParseForm()
	pages := ctx.Request.Form.Get(conf.REWARD_PAGES)
	pagesInt, _ := strconv.Atoi(pages)
	rewards := models.ShowReward(pagesInt)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func DeleteReward(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.REWARD_PUBLISH_USER_ID)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	models.DeleteReward(rewardId, userId)
	ctx.Output.JSON(models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}, true, false)
}

func ShowRewardSortXiaodian(ctx *context.Context) {
	ctx.Request.ParseForm()
	pages := ctx.Request.Form.Get(conf.REWARD_PAGES)
	pagesInt, _ := strconv.Atoi(pages)
	rewards := models.ShowRewardSortXiaodian(pagesInt)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func ShowRewardKeyword(ctx *context.Context) {
	ctx.Request.ParseForm()
	pages := ctx.Request.Form.Get(conf.REWARD_PAGES)
	keyword := ctx.Request.Form.Get(conf.KEYWORD)
	pagesInt, _ := strconv.Atoi(pages)
	rewards := models.ShowRewardKeyword(pagesInt, keyword)
	rewardResp := models.RewardResp{conf.SUCCESS, conf.SUCCESS_MSG, rewards}
	ctx.Output.JSON(rewardResp, true, false)
}

func CarryReward(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	models.CarryReward(rewardId, userId)
	baseResp := models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}
	ctx.Output.JSON(baseResp, true, false)
}

func DeliveryReward(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	models.DeliveryReward(rewardId, userId)
	baseResp := models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}
	ctx.Output.JSON(baseResp, true, false)
}

func FinishReward(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	models.FinishReward(rewardId, userId)
	baseResp := models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}
	ctx.Output.JSON(baseResp, true, false)
}

func Evaluate(ctx *context.Context) {
	ctx.Request.ParseForm()
	userId := ctx.Request.Form.Get(conf.TOKEN_USER_ID)
	rewardId := ctx.Request.Form.Get(conf.REWARD_ID)
	evaluate := ctx.Request.Form.Get(conf.EVALUATE)

	evaluateFloat, err := strconv.ParseFloat(evaluate, 32)
	if nil != err {
		panic(err)
	}
	models.EvaluateReward(rewardId, userId, float32(evaluateFloat))
	baseResp := models.BaseResp{conf.SUCCESS, conf.SUCCESS_MSG}
	ctx.Output.JSON(baseResp, true, false)
}

func getThing(ctx *context.Context) models.Thing {
	thingType := ctx.Request.Form.Get(conf.THING_TYPE)
	thingTypeInt, err := strconv.ParseInt(thingType, 10, 32)
	if nil != err {
		panic(models.BaseResp{conf.ERROR_THING_TYPE, conf.ERROR_THING_TYPE_MSG})
	}
	thumbnail := ""
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
