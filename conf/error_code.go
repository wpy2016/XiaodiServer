package conf

const (
	SUCCESS     = 200
	SUCCESS_MSG = "success"

	FAIL = 444
	//error
	ERROR_IMG_KIND_TYPE     = 3001
	ERROR_IMG_KIND_TYPE_MSG = "图片类型错误"

	ERROR_PHONE_IS_EXIST     = 3002
	ERROR_PHONE_IS_EXIST_MSG = "手机号已经注册"

	ERROR_NICKNAME_EXIST     = 3003
	ERROR_NICKNAME_EXIST_MSG = "昵称已经使用"

	ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR     = 3004
	ERROR_ACCOUNT_NOT_EXIST_OR_PASS_ERROR_MSG = "手机号或密码错误"

	ERROR_NOT_HAVE_TOKEN     = 3005
	ERROR_NOT_HAVE_TOKEN_MSG = "不存在Token"

	ERROR_USER_NOT_EXIST     = 3006
	ERROR_USER_NOT_EXIST_MSG = "用户不存在"

	ERROR_USER_ALREADY_AUTH     = 2001
	ERROR_USER_ALREADY_AUTH_MSG = "用户已经认证"

	ERROR_USER_XIAODIAN_SHORT     = 2003
	ERROR_USER_XIAODIAN_SHORT_MSG = "笑点不足"

	ERROR_TIME_LAYOUT     = 3007
	ERROR_TIME_LAYOUT_MSG = "时间格式错误"

	ERROR_THING_TYPE     = 3008
	ERROR_THING_TYPE_MSG = "物品类型错误"

	ERROR_XIAODIAN_LAYOUT     = 3009
	ERROR_XIAODIAN_LAYOUT_MSG = "笑点格式错误"

	REWARD_CAN_NOT_CARRY     = 3010
	REWARD_CAN_NOT_CARRY_MSG = "悬赏已经被领取或已经过期"

	REWARD_CARRY_NEED_PERMISSION     = 3011
	REWARD_CARRY_NEED_PERMISSION_MSG = "要认证成为笑递员才可以领取哦"

	REWARD_NOT_EXIST     = 3012
	REWARD_NOT_EXIST_MSG = "悬赏不存在"

	NOT_SEND_REWARD_CAN_NOT_DELETE     = 3013
	NOT_SEND_REWARD_CAN_NOT_DELETE_MSG = "已领取悬赏不允许删除"

	NOT_OWNER_REWARD_CAN_NOT_DELETE     = 3014
	NOT_OWNER_REWARD_CAN_NOT_DELETE_MSG = "不允许删除非本人的悬赏"

	NOT_SEND_REWARD_CAN_NOT_UPDATE     = 3015
	NOT_SEND_REWARD_CAN_NOT_UPDATE_MSG = "已领取悬赏不允许更新"

	NOT_OWNER_REWARD_CAN_NOT_UPDATE     = 3016
	NOT_OWNER_REWARD_CAN_NOT_UPDATE_MSG = "不允许更新非本人的悬赏"

	HAVE_NOT_PERMISSION     = 3017
	HAVE_NOT_PERMISSION_MSG = "你没有权限"

	REWARD_NOT_FINISH     = 3018
	REWARD_NOT_FINISH_MSG = "此悬赏还未完成"

	REWARD_ALREADY_EVALUATE     = 3019
	REWARD_ALREADY_EVALUATE_MSG = "已经评价了"

	TODAY_IS_SIGN     = 3020
	TODAY_IS_SIGN_MSG = "今天已经签到了"

	THIS_MONTH_YOU_HAVE_NOT_SIGN     = 3021
	THIS_MONTH_YOU_HAVE_NOT_SIGN_MSG = "这个月暂无签到信息"

	OLD_PASS_ERROR     = 3022
	OLD_PASS_ERROR_MSG = "旧密码错误"
)
