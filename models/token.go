package models

import (
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserToken struct {
	UserId string `json:"user_id" bson:"user_id"`
	Token  string `json:"token" bson:"token"`
}

func (token *UserToken) Save() error {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	tokenC := session.DB(conf.MGO_DB).C(conf.MGO_DB_TOKEN_COLLECTION)
	err = tokenC.Insert(token)
	return err
}

func AssertTokenExist(userId, token string) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	tokenC := session.DB(conf.MGO_DB).C(conf.MGO_DB_TOKEN_COLLECTION)
	userToken := &UserToken{}
	err = tokenC.Find(bson.M{conf.TOKEN_USER_ID: userId, conf.TOEKN: token}).One(userToken)
	if nil != err {
		panic(BaseResp{conf.ERROR_NOT_HAVE_TOKEN, conf.ERROR_NOT_HAVE_TOKEN_MSG})
	}
}
