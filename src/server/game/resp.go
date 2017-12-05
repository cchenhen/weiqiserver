package game

type RESP_Weiqi_01 struct {
	Status       uint32
	OnlinePlayer []string
}

type RESP_Weiqi_02 struct {
	Status       uint32
	OnlinePlayer []string
	AllGameInfo  []AllGameInfo
}

type RESP_Weiqi_03 struct {
	Status uint32
	GameId int64
}

type RESP_Weiqi_04 struct {
	Status     uint32
	GameStatus []int64
}

type RESP_Weiqi_06 struct {
	Status     uint32
	Round      uint32
	Player     [2]string
	Size       uint32
	GameStatus []int64
}

type AllGameInfo struct {
	GameId     string
	Round      uint32
	Player     [2]string
	Size       uint32
	GameStatus []int64
}
