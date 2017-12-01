package game

type RESP_Weiqi_01 struct {
	Status       uint32
	OnlinePlayer []string
}

type RESP_Weiqi_02 struct {
	Status       uint32
	LiveGame     []int64
	OnlinePlayer []string
}

type RESP_Weiqi_03 struct {
	Status uint32
	GameId int64
}
