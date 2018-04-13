package models

import (
	"gopkg.in/mgo.v2"
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

/**
通过短信验证码进行修改密码使用的认证方式
 */
type OneToken struct {
	Id    string `json:"_id" bson:"_id"`
	Phone string `json:"phone" bson:"phone"`
	Token string `json:"token" bson:"token"`
}

type OneTokenResp struct {
	StatusCode int    `json:"status_code" bson:"status_code"`
	StatusMsg  string `json:"status_msg" bson:"status_msg"`
	Token      string `json:"token" bson:"token"`
}

func AuthOneToken(phone, token string) {
	session, tokenC := getOneTokenDbCollection()
	defer session.Close()
	fmt.Println(phone,"        ",token)
	oneToken := &OneToken{}
	err := tokenC.Find(bson.M{conf.USER_PHONE: phone, conf.TOEKN: token}).One(oneToken)
	if nil != err {
		panic(BaseResp{conf.ERROR_NOT_HAVE_TOKEN, conf.ERROR_NOT_HAVE_TOKEN_MSG})
	}
	err = tokenC.Remove(bson.M{conf.ID: oneToken.Id})
	if nil != err {
		fmt.Println(err)
		panic(err)
	}
}

func GetOneToken(phone string) string {
	session, oneTokenC := getOneTokenDbCollection()
	defer session.Close()
	token := bson.NewObjectId().Hex()
	oneToken := &OneToken{
		Id:    bson.NewObjectId().Hex(),
		Phone: phone,
		Token: token,
	}
	err := oneTokenC.Insert(oneToken)
	if nil != err {
		panic(err)
	}
	return token
}

func getOneTokenDbCollection() (*mgo.Session, *mgo.Collection) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	oneTokenC := session.DB(conf.MGO_DB).C(conf.MGO_DB_ONE_TOKEN_COLLECTION)
	return session, oneTokenC
}
