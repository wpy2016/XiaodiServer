package models

import (
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2"
	"time"
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
}

func (reward *Reward) Save() {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	rewardC := session.DB(conf.MGO_DB).C(conf.MGO_DB_REWARD_COLLECTION)
	err = rewardC.Insert(reward)
	if nil != err {
		panic(err)
	}
}
