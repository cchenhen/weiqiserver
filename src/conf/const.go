package conf

//基础信息
const (
	BLACK_PLAYER    = 0 // 黑子玩家
	WHITE_PLAYER    = 1 // 白棋玩家
	WINNER          = 1 // 胜利者
	FAILER          = 2 // 失败者
	UN_FINISHED     = 0 // 未完成对局
	FINISHED_WIN    = 1 // 完成并且胜利
	FINISHED_FAILED = 2 // 完成并且失败

	GIVE_UP = 400 // 放弃该手
)

//错误码
const (
	SUCCEED            = 0
	ERR_BAD_PARAM      = 100 //错误滴参数
	ERR_SERVER_ERR     = 113 //服务器错误
	ERR_INVITE_OFFLINE = 101 //邀请好友不在线
)

// 棋盘型号
const (
	WEIQI_SIZE_SMALL    = 9
	WEIQI_SIZE_MID      = 13
	WEIQI_SIZE_STANDARD = 19
)

//棋盘信息
const (
	SPACE_G = 0 //二进制为00
	BLACK_G = 2 //二进制为10
	WHITE_G = 3 //二进制为11
)
