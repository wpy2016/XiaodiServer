package models

import (
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
)

type Thing struct {
	ThingType int    `json:"thing_type" bson:"thing_type"` //0表示快递，1表示餐饮，2表示纸质，3表示其他
	Thumbnail string `json:"thumbnail" bson:"thumbnail"`
	Weight    string `json:"weight" bson:"weight"`
}

type Reward struct {
	ID             string    `json:"_id" bson:"_id"`
	Publisher      BaseUser  `json:"publisher" bson:"publisher"`
	State          int       `json:"state" bson:"state"` //-1 表示过期,0表示发布,1表示拿到了物品并代送中,2表示笑递员发起带到，3表示物主完成
	Phone          string    `json:"phone" bson:"phone"`
	Xiaodian       int       `json:"xiaodian" bson:"xiaodian"`
	DeadLine       time.Time `json:"dead_line" bson:"dead_line"`
	OriginLocation string    `json:"origin_location" bson:"origin_location"`
	DstLocation    string    `json:"dst_location" bson:"dst_location"`
	Receiver       BaseUser  `json:"receiver" bson:"receiver"`
	PublisherGrade float32   `json:"publisher_grade" bson:"publisher_grade"`
	ReceiveGrade   float32   `json:"receive_grade" bson:"receive_grade"`
	Describe       string    `json:"describe" bson:"describe"`
	Thing          Thing     `json:"thing" bson:"thing"`
	CreateTime     time.Time `json:"create_time" bson:"create_time"`
}

func (reward *Reward) Save() {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	err := rewardC.Insert(reward)
	if nil != err {
		panic(err)
	}
}

func ShowReward(pages int) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	timeNow := time.Now()
	timeStr := timeNow.Format(conf.TIME_FORMAT)
	formatTime, _ := time.Parse(conf.TIME_FORMAT, timeStr)
	var rewards []Reward
	err := rewardC.Find(bson.M{"dead_line": bson.M{"$gt": formatTime}, "state": conf.REWARD_SEND}).Sort("-create_time").Limit(
		conf.REWARD_PAGES_ITEM_COUNT).Skip(pages * conf.REWARD_PAGES_ITEM_COUNT).All(&rewards)
	if nil != err {
		panic(errors.New("ShowReward" + err.Error()))
	}
	return rewards
}

func CarryReward(rewardId,userId string) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	user := GetBaseUserById(userId)
	if conf.NORMAL_USER == user.UserType {
		panic(BaseResp{conf.REWARD_CARRY_NEED_PERMISSION,conf.REWARD_CARRY_NEED_PERMISSION_MSG})
	}
	timeNow := time.Now()
	timeStr := timeNow.Format(conf.TIME_FORMAT)
	formatTime, _ := time.Parse(conf.TIME_FORMAT, timeStr)
	reward := &Reward{}
	err := rewardC.Find(bson.M{conf.REWARD_DEADLINE: bson.M{"$gt": formatTime},conf.REWARD_STATE: conf.REWARD_SEND,conf.ID:rewardId}).One(reward)
	if nil != err {
		panic(BaseResp{conf.REWARD_CAN_NOT_CARRY,conf.REWARD_CAN_NOT_CARRY_MSG})
	}
	err = rewardC.Update(bson.M{conf.ID: rewardId}, bson.M{
		"$set": bson.M{
			"state": conf.REWARD_CARRY,
			"receiver":user,
		}})
	if nil != err {
		panic(BaseResp{444,err.Error()})
	}
}

func getRewardDbCollection() (*mgo.Session, *mgo.Collection) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	rewardC := session.DB(conf.MGO_DB).C(conf.MGO_DB_REWARD_COLLECTION)
	return session, rewardC
}