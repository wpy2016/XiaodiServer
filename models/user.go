package models

type User struct {
	ID         string `json:"_id",bson:"_id"`
	NickName   string `json:"nick_name",bson:"nick_name"`
	RealName   string `json:"real_name",bson:"real_name"`
	Phone      string `json:"phone",bson:"phone"`
	Pass       string `json:"pass",bson:"pass"`
	CreateTime string `json:"create_time",bson:"create_time"`
	UserType   string `json:"user_type",bson:"user_type"`
	Img        string `json:"img",bson:"img"`
}

type LinkedUser func() error

func (user *User) Save() error {

	return nil
}

func (user *User) IsPhoneUse(fun *LinkedUser) (bool, error) {

	return false, nil
}

func (user *User) isNickNameUse(fun *LinkedUser) (bool, error) {

	return false, nil
}

func (user *User) Update() error {

	return nil
}

func (user *User) Login() error {
	return nil
}
