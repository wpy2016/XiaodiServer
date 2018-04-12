package models

import (
	"gopkg.in/mgo.v2/bson"
	"XiaodiServer/conf"
	"gopkg.in/mgo.v2"
	"XiaodiServer/models/db"
	"time"
	"strconv"
)

type Sign struct {
	Id     string   `json:"_id" bson:"_id"`
	UserId string   `json:"user_id" bson:"user_id"`
	Year   string   `json:"year" bson:"year"`
	Month  string   `json:"month" bson:"month"`
	Days   []string `json:"days" bson:"days"`
}

func (sign *Sign) Save() {
	session, signC := getSignDbCollection()
	defer session.Close()
	err := signC.Insert(sign)
	if nil != err {
		panic(err)
	}
}

func SignToday(userId, day string) {
	session, signC := getSignDbCollection()
	defer session.Close()
	sign := &Sign{}
	year := strconv.Itoa(time.Now().Year())
	month := strconv.Itoa(int(time.Now().Month()))
	err := signC.Find(bson.M{conf.TOKEN_USER_ID: userId, conf.SIGN_YEAR: year, conf.SIGN_MONTH: month}).One(sign)
	if nil != err {
		newSign := &Sign{
			Id:     bson.NewObjectId().Hex(),
			UserId: userId,
			Year:   year,
			Month:  month,
			Days:   []string{day},
		}
		newSign.Save()
		return
	}
	if isSign(sign.Days, day) {
		panic(BaseResp{StatusCode: conf.TODAY_IS_SIGN, StatusMsg: conf.TODAY_IS_SIGN_MSG})
	}
	sign.Days = append(sign.Days, day)
	err = signC.Update(bson.M{conf.ID: sign.Id}, sign)
	if nil != err {
		panic(err)
	}
}

func isSign(days []string, day string) bool {
	for _, d := range days {
		if d == day {
			return true
		}
	}
	return false
}

func MySignList(userId, year, month string) []string {
	session, signC := getSignDbCollection()
	defer session.Close()
	sign := &Sign{}
	err := signC.Find(bson.M{conf.TOKEN_USER_ID: userId, conf.SIGN_YEAR: year, conf.SIGN_MONTH: month}).One(sign)
	if nil != err {
		panic(BaseResp{StatusCode: conf.THIS_MONTH_YOU_HAVE_NOT_SIGN, StatusMsg: conf.THIS_MONTH_YOU_HAVE_NOT_SIGN_MSG})
	}
	return sign.Days
}

func getSignDbCollection() (*mgo.Session, *mgo.Collection) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	signC := session.DB(conf.MGO_DB).C(conf.MGO_DB_SIGN_COLLECTION)
	return session, signC
}
