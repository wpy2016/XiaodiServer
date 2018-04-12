package models

type SignResp struct {
	StatusCode int      `json:"status_code" bson:"status_code"`
	StatusMsg  string   `json:"status_msg" bson:"status_msg"`
	Days       []string `json:"days" bson:"days"`
}
