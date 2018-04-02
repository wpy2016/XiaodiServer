package db

import (
	"XiaodiServer/conf"
	"gopkg.in/mgo.v2"
	"time"
)

func CreateDialInfo() *mgo.DialInfo {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{conf.MGO_IP},
		Direct:    false,
		Timeout:   time.Second * 2,
		Source:    conf.MGO_DB,
		Database:  conf.MGO_DB,
		Username:  conf.MGO_LOGIN_USER,
		Password:  conf.MGO_LOGIN_PASS,
		PoolLimit: 4096, // Session.SetPoolLimit
	}
	return dialInfo
}
