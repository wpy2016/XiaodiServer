package models

import (
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"errors"
	"fmt"
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
	DeadLine       string    `json:"dead_line" bson:"dead_line"`
	DeadLineTime   time.Time `json:"dead_line_time" bson:"dead_line_time"`
	OriginLocation string    `json:"origin_location" bson:"origin_location"`
	DstLocation    string    `json:"dst_location" bson:"dst_location"`
	Receiver       BaseUser  `json:"receiver" bson:"receiver"`
	PublisherGrade int       `json:"publisher_grade" bson:"publisher_grade"`
	ReceiveGrade   int       `json:"receive_grade" bson:"receive_grade"`
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

func UpdateReward(rewardId, userId string, newReward Reward) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	var oldReward Reward
	err := rewardC.Find(bson.M{conf.ID: rewardId}).One(&oldReward)
	if nil != err {
		panic(BaseResp{StatusCode: conf.REWARD_NOT_EXIST, StatusMsg: conf.REWARD_NOT_EXIST_MSG})
	}
	if userId != oldReward.Publisher.ID {
		panic(BaseResp{StatusCode: conf.NOT_OWNER_REWARD_CAN_NOT_UPDATE, StatusMsg: conf.NOT_OWNER_REWARD_CAN_NOT_UPDATE_MSG})
	}
	if conf.REWARD_SEND != oldReward.State {
		panic(BaseResp{StatusCode: conf.NOT_SEND_REWARD_CAN_NOT_UPDATE, StatusMsg: conf.NOT_SEND_REWARD_CAN_NOT_UPDATE_MSG})
	}
	err = rewardC.Update(bson.M{conf.ID: oldReward.ID}, newReward)
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
	err := rewardC.Find(bson.M{conf.REWARD_DEADLINE_TIME: bson.M{"$gt": formatTime}, "state": conf.REWARD_SEND}).Sort("-create_time").Limit(
		conf.REWARD_PAGES_ITEM_COUNT).Skip(pages * conf.REWARD_PAGES_ITEM_COUNT).All(&rewards)
	if nil != err {
		panic(errors.New("ShowReward" + err.Error()))
	}
	return rewards
}

func DeleteReward(id, userId string) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	var reward Reward
	err := rewardC.Find(bson.M{conf.ID: id}).One(&reward)
	if nil != err {
		panic(BaseResp{StatusCode: conf.REWARD_NOT_EXIST, StatusMsg: conf.REWARD_NOT_EXIST_MSG})
	}
	if userId != reward.Publisher.ID {
		panic(BaseResp{StatusCode: conf.NOT_OWNER_REWARD_CAN_NOT_DELETE, StatusMsg: conf.NOT_OWNER_REWARD_CAN_NOT_DELETE_MSG})
	}
	if conf.REWARD_SEND != reward.State {
		panic(BaseResp{StatusCode: conf.NOT_SEND_REWARD_CAN_NOT_DELETE, StatusMsg: conf.NOT_SEND_REWARD_CAN_NOT_DELETE_MSG})
	}
	err = rewardC.Remove(bson.M{conf.ID: id})
	if nil != err {
		panic(err)
	}
}

func ShowRewardMySend(userId string) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	var rewards []Reward
	err := rewardC.Find(bson.M{"publisher._id": userId}).Sort("-create_time").All(&rewards)
	if nil != err {
		panic(errors.New("ShowRewardMySend" + err.Error()))
	}
	return rewards
}

func ShowRewardOurNotFinish(userId, receiveId string) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	var rewardmySend []Reward
	err := rewardC.Find(bson.M{"publisher._id": userId, "receiver._id": receiveId, "state": bson.M{"$ne": conf.REWARD_FINISH}}).Sort("-create_time").All(&rewardmySend)
	if nil != err {
		panic(errors.New("ShowRewardOur" + err.Error()))
	}
	var rewardmyCarry []Reward
	err = rewardC.Find(bson.M{"publisher._id": receiveId, "receiver._id": userId, "state": bson.M{"$ne": conf.REWARD_FINISH}}).Sort("-create_time").All(&rewardmyCarry)
	if nil != err {
		panic(errors.New("ShowRewardOur" + err.Error()))
	}

	var rewards []Reward
	rewards = append(rewards, rewardmySend...)
	rewards = append(rewards, rewardmyCarry...)
	return rewards
}

func ShowRewardMyCarry(userId string) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	var rewards []Reward
	err := rewardC.Find(bson.M{"receiver._id": userId}).Sort("-create_time").All(&rewards)
	if nil != err {
		panic(errors.New("ShowRewardMyCarry" + err.Error()))
	}
	return rewards
}

func ShowRewardSortXiaodian(pages int) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	timeNow := time.Now()
	timeStr := timeNow.Format(conf.TIME_FORMAT)
	formatTime, _ := time.Parse(conf.TIME_FORMAT, timeStr)
	var rewards []Reward
	err := rewardC.Find(bson.M{conf.REWARD_DEADLINE_TIME: bson.M{"$gt": formatTime}, "state": conf.REWARD_SEND}).Sort("-xiaodian").Limit(
		conf.REWARD_PAGES_ITEM_COUNT).Skip(pages * conf.REWARD_PAGES_ITEM_COUNT).All(&rewards)
	if nil != err {
		panic(errors.New("ShowReward" + err.Error()))
	}
	return rewards
}

func ShowRewardKeyword(pages int, key string) []Reward {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	timeNow := time.Now()
	timeStr := timeNow.Format(conf.TIME_FORMAT)
	formatTime, _ := time.Parse(conf.TIME_FORMAT, timeStr)
	var rewards []Reward
	err := rewardC.Find(bson.M{
		conf.REWARD_DEADLINE_TIME: bson.M{"$gt": formatTime},
		"state":                   conf.REWARD_SEND,
		//	"$text":bson.M{"$search":key},
	}).Sort("-xiaodian").Limit(
		conf.REWARD_PAGES_ITEM_COUNT).Skip(pages * conf.REWARD_PAGES_ITEM_COUNT).All(&rewards)
	if nil != err {
		panic(errors.New("ShowReward" + err.Error()))
	}
	return rewards
}

func CarryReward(rewardId, userId string) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	user := GetBaseUserById(userId)
	if conf.NORMAL_USER == user.UserType {
		panic(BaseResp{conf.REWARD_CARRY_NEED_PERMISSION, conf.REWARD_CARRY_NEED_PERMISSION_MSG})
	}
	timeNow := time.Now()
	timeStr := timeNow.Format(conf.TIME_FORMAT)
	formatTime, _ := time.Parse(conf.TIME_FORMAT, timeStr)
	reward := &Reward{}
	err := rewardC.Find(bson.M{conf.REWARD_DEADLINE_TIME: bson.M{"$gt": formatTime}, conf.REWARD_STATE: conf.REWARD_SEND, conf.ID: rewardId}).One(reward)
	if nil != err {
		panic(BaseResp{conf.REWARD_CAN_NOT_CARRY, conf.REWARD_CAN_NOT_CARRY_MSG})
	}
	err = rewardC.Update(bson.M{conf.ID: rewardId}, bson.M{
		"$set": bson.M{
			"state":    conf.REWARD_CARRY,
			"receiver": user,
		}})
	if nil != err {
		panic(BaseResp{444, err.Error()})
	}
}

func DeliveryReward(rewardId, userId string) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	reward := &Reward{}
	err := rewardC.Find(bson.M{conf.ID: rewardId}).One(reward)
	if nil != err {
		panic(BaseResp{conf.REWARD_NOT_EXIST, conf.REWARD_NOT_EXIST_MSG})
	}
	if (userId != reward.Receiver.ID) {
		panic(BaseResp{conf.HAVE_NOT_PERMISSION, conf.HAVE_NOT_PERMISSION_MSG})
	}
	err = rewardC.Update(bson.M{conf.ID: rewardId}, bson.M{
		"$set": bson.M{
			"state": conf.REWARD_ARRIVE,
		}})
	if nil != err {
		panic(BaseResp{444, err.Error()})
	}
}

func FinishReward(rewardId, userId string) {
	session, rewardC := getRewardDbCollection()
	defer session.Close()
	reward := &Reward{}
	err := rewardC.Find(bson.M{conf.ID: rewardId}).One(reward)
	if nil != err {
		panic(BaseResp{conf.REWARD_NOT_EXIST, conf.REWARD_NOT_EXIST_MSG})
	}
	if (userId != reward.Publisher.ID) {
		panic(BaseResp{conf.HAVE_NOT_PERMISSION, conf.HAVE_NOT_PERMISSION_MSG})
	}
	//进行笑点转移
	moveXiaodian(reward.Publisher.ID, reward.Receiver.ID, reward.Xiaodian)

	err = rewardC.Update(bson.M{conf.ID: rewardId}, bson.M{
		"$set": bson.M{
			"state": conf.REWARD_FINISH,
		}})
	if nil != err {
		panic(BaseResp{444, err.Error()})
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
