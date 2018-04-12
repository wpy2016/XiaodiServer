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
	GoldMoney     int     `json:"gold_money" bson:"gold_money"`       //金笑点
	SilverMoney   int     `json:"silver_money" bson:"silver_money"`   //银笑点
	Sign          string  `json:"sign" bson:"sign"`                   //签名
	Token         string  `json:"token" bson:"token"`                 //身份认证，需要不定时更新,这个token用于本服务器
	RongyunToken  string  `json:"rongyun_token" bson:"rongyun_token"` //融云IM即时通讯的token，每个用户的标识
}

func (user *User) Save() {
	session, userC := getUserDbCollection()
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
	session, userC := getUserDbCollection()
	defer session.Close()
	user := &User{}
	err := userC.Find(bson.M{conf.USER_PHONE: phone, conf.USER_PASS: pass}).One(user)
	if nil != err {
		panic(BaseResp{conf.ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR, conf.ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR_MSG})
	}
	return user
}

func Auth(userId, realName, schoolId, campus string) {
	session, userC := getUserDbCollection()
	defer session.Close()
	user := &User{}
	err := userC.Find(bson.M{conf.ID: userId}).One(user)
	if nil != err {
		panic(BaseResp{StatusCode: conf.ERROR_USER_NOT_EXIST, StatusMsg: conf.ERROR_USER_NOT_EXIST_MSG})
	}
	if conf.XIAODI_YUAN == user.UserType {
		panic(BaseResp{StatusCode: conf.ERROR_USER_ALREADY_AUTH, StatusMsg: conf.ERROR_USER_ALREADY_AUTH_MSG})
	}
	user.UserType = conf.XIAODI_YUAN
	user.Campus = campus
	user.RealName = realName
	user.SchoolID = schoolId
	err = userC.Update(bson.M{conf.ID: userId}, user)
	if nil != err {
		panic(err)
	}
}

func UpdatePass(userId, oldDecryptPass, newDecryptPss string) {
	user := GetUserById(userId)
	if oldDecryptPass != user.Pass {
		panic(BaseResp{conf.OLD_PASS_ERROR, conf.OLD_PASS_ERROR_MSG})
	}
	session, userC := getUserDbCollection()
	defer session.Close()
	err := userC.Update(bson.M{conf.ID: user.ID}, bson.M{
		"$set": bson.M{
			"pass": newDecryptPss,
		}})
	if nil != err {
		panic(err)
	}
}

func UpdateNickname(userId, nickName string) {
	user := GetUserById(userId)
	session, userC := getUserDbCollection()
	defer session.Close()
	//判断用户名是否已经存在
	userToCheckNickName := &User{}
	err := userC.Find(bson.M{conf.USER_NICKNAME: nickName}).One(userToCheckNickName)
	if nil == err {
		panic(BaseResp{conf.ERROR_NICKNAME_EXIST, conf.ERROR_NICKNAME_EXIST_MSG})
	}

	err = nil
	err = userC.Update(bson.M{conf.ID: user.ID}, bson.M{
		"$set": bson.M{
			"nick_name": nickName,
		}})
	if nil != err {
		panic(err)
	}
}

func UpdateImg(userId, imgurl string) {
	user := GetUserById(userId)
	session, userC := getUserDbCollection()
	defer session.Close()
	err := userC.Update(bson.M{conf.ID: user.ID}, bson.M{
		"$set": bson.M{
			"img": imgurl,
		}})
	if nil != err {
		panic(err)
	}
}

func IsPhoneExist(phone string) (bool, error) {
	session, userC := getUserDbCollection()
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
	session, userC := getUserDbCollection()
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
	session, userC := getUserDbCollection()
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

func getUserDbCollection() (*mgo.Session, *mgo.Collection) {
	dialInfo := db.CreateDialInfo()
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	userC := session.DB(conf.MGO_DB).C(conf.MGO_DB_USER_COLLECTION)
	return session, userC
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

func moveXiaodian(publisherId, receiverId string, xiaodian int) {
	publisher := GetUserById(publisherId)
	receiver := GetUserById(receiverId)
	//首先使用充值进来的笑点进行消费,再使用代送获取的笑点，都不足，便一起使用
	if publisher.GoldMoney >= xiaodian {
		publisher.GoldMoney = publisher.GoldMoney - xiaodian
		receiver.SilverMoney = receiver.SilverMoney + xiaodian
		updateUser(publisher)
		updateUser(receiver)
		return
	}

	//金笑点不够，但金银笑点加起来够,这时候，消耗全部金笑点，银笑点再减去不足的部分
	if publisher.GoldMoney+publisher.SilverMoney >= xiaodian {
		publisher.SilverMoney = publisher.SilverMoney - (xiaodian - publisher.GoldMoney)
		publisher.GoldMoney = 0
		receiver.SilverMoney = receiver.SilverMoney + xiaodian
		updateUser(publisher)
		updateUser(receiver)
		return
	}

	//发布者笑点不足
	panic(BaseResp{conf.ERROR_USER_XIAODIAN_SHORT, conf.ERROR_USER_XIAODIAN_SHORT_MSG})
}

func updateUser(user *User) {
	GetUserById(user.ID) //主要用来检查错误，用户是否存在
	session, userC := getUserDbCollection()
	defer session.Close()
	err := userC.Update(bson.M{conf.ID: user.ID}, user)
	if nil != err {
		panic(err)
	}
}
