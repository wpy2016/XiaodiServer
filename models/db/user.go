package db

import (
	"XiaodiServer/conf"
	"XiaodiServer/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Save(user *models.User) error {
	dialInfo := CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	userC := session.DB(conf.MGO_DB).C(conf.MGO_DB_USER_COLLECTION)
	user.ID = bson.NewObjectId().Hex()
	err = userC.Insert(user)
	return err
}

func Update(user *models.User) error {

	return nil
}

func isPhoneExist(phone string) (bool, error) {
	dialInfo := CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		return false, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	userC := session.DB(conf.MGO_DB).C(conf.MGO_DB_USER_COLLECTION)
	count, err := userC.Find(bson.M{conf.USER_PHONE: phone}).Count()
	if nil != err {
		return false, err
	}
	if 0 != count {
		return true, nil
	}
	return false, nil
}

func isNickNameExist(nickName string) (bool, error) {
	dialInfo := CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if nil != err {
		return false, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	userC := session.DB(conf.MGO_DB).C(conf.MGO_DB_USER_COLLECTION)
	count, err := userC.Find(bson.M{conf.USER_NICKNAME: nickName}).Count()
	if nil != err {
		return false, err
	}
	if 0 != count {
		return true, nil
	}
	return false, nil
}
