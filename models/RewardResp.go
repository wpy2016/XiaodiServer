package models

type RewardResp struct {
	StatusCode int      `json:"status_code" bson:"status_code"`
	StatusMsg  string   `json:"status_msg" bson:"status_msg"`
	Rewards    []Reward `json:"rewards" bson:"rewards"`
}
