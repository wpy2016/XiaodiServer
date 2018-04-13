package conf

const (
	UPLOAD_IMG_HEAD_FILE_PATH   = "Xiaodi/ImgHead"
	UPLOAD_IMG_REWARD_FILE_PATH = "Xiaodi/reward"
	IMG_HEAD_HTTP               = "http://119.29.164.153:1688/picture/0/"
	IMG_REWARD_HTTP             = "http://119.29.164.153:1688/picture/1/"

	NORMAL_USER          = 0
	XIAODI_YUAN          = 1
	DEFAULt_SLIVER_MONEY = 300

	THING_TYPE_EXPRESS = 0
	THING_TYPE_FOOD    = 1
	THING_TYPE_PAPER   = 2
	THING_TYPE_OTHER   = 3

	REWARD_EXPIRED = -1
	REWARD_SEND    = 0
	REWARD_CARRY   = 1
	REWARD_ARRIVE  = 2
	REWARD_FINISH  = 3

	DEFAULT_THING_EXPRESS_IMG = IMG_REWARD_HTTP + "express.png"
	DEFAULT_THING_FOOD_IMG    = IMG_REWARD_HTTP + "food.png"
	DEFAULT_THING_PAPER_IMG   = IMG_REWARD_HTTP + "paper.png"
	DEFAULT_THING_OTHER_IMG   = IMG_REWARD_HTTP + "other.png"

	SERVICE_PHONE = "15111929296"
	XIAODIAN_CENTER = "笑点中心"

	SIGN_ADD_XIAODIAN = "签到增加1笑点"
	XIAODIAN_THUMBNAIL = "http://119.29.164.153:1688/picture/0/xiaodi.png"
)
