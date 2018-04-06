package models

type UserResp struct {
	StatusCode int    `json:"status_code" bson:"status_code"`
	StatusMsg  string `json:"status_msg" bson:"status_msg"`
	User       User   `json:"user" bson:"user"`
}
