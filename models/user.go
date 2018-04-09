package models

import (
	"XiaodiServer/conf"
	"XiaodiServer/models/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BaseUser struct {
	ID            string  `json:"_id" bson:"_id"`
	NickName      string  `json:"nick_name" bson:"nick_name"`         //昵称
	RealName      string  `json:"real_name" bson:"real_name"`         //真实姓名
	Phone         string  `json:"phone" bson:"phone"`                 //手机号
	UserType      int     `json:"user_type" bson:"user_type"`         //0表示普通用户，1表示笑递员
	Campus        string  `json:"campus" bson:"campus"`               //学生所在学院
	SchoolID      string  `json:"school_id" bson:"school_id"`         //学号
	Img           string  `json:"img" bson:"img"`                     //头像
	Creditibility float32 `json:"creditibility" bson:"creditibility"` //信誉度
}
type User struct {
	ID            string  `json:"_id" bson:"_id"`
	NickName      string  `json:"nick_name" bson:"nick_name"`         //昵称
	RealName      string  `json:"real_name" bson:"real_name"`         //真实姓名
	Phone         string  `json:"phone" bson:"phone"`                 //手机号
	UserType      int     `json:"user_type" bson:"user_type"`         //0表示普通用户，1表示笑递员
	Campus        string  `json:"campus" bson:"campus"`               //学生所在学院
	SchoolID      string  `json:"school_id" bson:"school_id"`         //学号
	Img           string  `json:"img" bson:"img"`                     //头像
	Creditibility float32 `json:"creditibility" bson:"creditibility"` //信誉度
	Pass          string  `json:"pass" bson:"pass"`                   //密码
	GoldMoney     float32 `json:"gold_money" bson:"gold_money"`       //金笑点
	SilverMoney   float32 `json:"silver_money" bson:"silver_money"`   //银笑点
	Sign          string  `json:"sign" bson:"sign"`                   //签名
	Token         string  `json:"token" bson:"token"`                 //身份认证，需要不定时更新
}

func (user *User) Save() {
	session,userC := getUserDbCollection()
	defer session.Close()
	err := userC.Insert(user)
	if nil != err {
		panic(err)
	}
}

//todo
func UpdateUser(user *User) error {

	return nil
}

func Login(phone, pass string) *User {
	session,userC := getUserDbCollection()
	defer session.Close()
	user := &User{}
	err := userC.Find(bson.M{conf.USER_PHONE: phone, conf.USER_PASS: pass}).One(user)
	if nil != err {
		panic(BaseResp{conf.ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR, conf.ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR_MSG})
	}
	return user
}
func IsPhoneExist(phone string) (bool, error) {
	session,userC := getUserDbCollection()
	defer session.Close()
	count, err := userC.Find(bson.M{conf.USER_PHONE: phone}).Count()
	if nil != err {
		return false, err
	}
	if 0 != count {
		return true, nil
	}
	return false, nil
}

func IsNickNameExist(nickName string) (bool, error) {
	session,userC := getUserDbCollection()
	defer session.Close()
	count, err := userC.Find(bson.M{conf.USER_NICKNAME: nickName}).Count()
	if nil != err {
		return false, err
	}
	if 0 != count {
		return true, nil
	}
	return false, nil
}

func GetUserById(id string) *User {
	session ,userC := getUserDbCollection()
	defer session.Close()
	user := &User{}
	err := userC.Find(bson.M{conf.ID: id}).One(user)
	if nil != err {
		panic(BaseResp{conf.ERROR_USER_NOT_EXIST, conf.ERROR_USER_NOT_EXIST_MSG})
	}
	return user
}

func GetBaseUserById(id string) *BaseUser {
	user := GetUserById(id)
	return &BaseUser{
		ID:            user.ID,
		NickName:      user.NickName,
		RealName:      user.RealName,
		Phone:         user.Phone,
		UserType:      user.UserType,
		Campus:        user.Campus,
		SchoolID:      user.SchoolID,
		Img:           user.Img,
		Creditibility: user.Creditibility,
	}
}

func getUserDbCollection() (*mgo.Session,*mgo.Collection) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	userC := session.DB(conf.MGO_DB).C(conf.MGO_DB_USER_COLLECTION)
	return session,userC
}

func RegisterDefaultUser(phone, decryptPass, nickName, imgPath string) *User {
	user := User{}
	user.ID = bson.NewObjectId().Hex()
	user.NickName = nickName
	user.UserType = conf.NORMAL_USER
	user.Creditibility = 5.0
	user.GoldMoney = 0
	user.Pass = decryptPass
	user.SilverMoney = conf.DEFAULt_SLIVER_MONEY
	user.Phone = phone
	user.Img = imgPath
	user.Token = bson.NewObjectId().Hex()
	return &user
}
