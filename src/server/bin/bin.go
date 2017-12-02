package bin

import (
	"conf"
	"log"
	"server/cache"
	"server/db"
	"server/game"
)

func Weiqi01(playerId string) *game.RESP_Weiqi_01 {
	player, err := db.GetPlayerInfo(playerId)
	if err != nil {
		log.Println("The first time to login in:%v", playerId)
		player = &game.PlayerInfo{}
		player.Default(playerId)
	}
	// add into PlayerList
	err = db.SetAllPlayerIdList(playerId)
	if err != nil {
		log.Println(err)
		return &game.RESP_Weiqi_01{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	// save playerinfo
	err = db.SetPlayerInfo(player.GetDbKey(), player)
	if err != nil {
		log.Println(err)
		return &game.RESP_Weiqi_01{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	cache.OutAddOnlinePlayer(player.PlayerId)
	onlineList := cache.GetAllOnlinePlayer(playerId)
	return &game.RESP_Weiqi_01{
		Status:       conf.SUCCEED,
		OnlinePlayer: onlineList,
	}
}

func Weiqi02(playerId string) *game.RESP_Weiqi_02 {
	player, err := db.GetPlayerInfo(playerId)
	if err != nil {
		log.Println("Never login", playerId)
		return &game.RESP_Weiqi_02{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	log.Println("PlayerInfo:", player)
	// keep alive
	err = db.SetAllPlayerIdList(playerId)
	if err != nil {
		log.Println(err)
		return &game.RESP_Weiqi_02{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	// get all onlineplayer
	cache.OutAddOnlinePlayer(player.PlayerId)
	onlineList := cache.GetAllOnlinePlayer(playerId)
	liveGame := player.GetOnGame()
	return &game.RESP_Weiqi_02{
		Status:       conf.SUCCEED,
		OnlinePlayer: onlineList,
		LiveGame:     liveGame,
	}
}

func Weiqi03(playerId string, inviteId string) *game.RESP_Weiqi_03 {
	player, err := db.GetPlayerInfo(playerId)
	if err != nil {
		log.Println("Never login", playerId)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	log.Println("PlayerInfo:", player)
	// check inviteid is alive
	if !cache.IsPlayerOnline(inviteId) {
		log.Println("Invite player is offline:", inviteId)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_INVITE_OFFLINE,
		}
	}
	invitePlayer, err := db.GetPlayerInfo(inviteId)
	if err != nil {
		log.Println("Never login", inviteId)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	// make a new game
	playerList := []string{playerId, inviteId}
	gameInfo := game.NewOneGame(playerList)
	// save to db
	err = db.SetRedisC(gameInfo.GetDbKey(), gameInfo)
	if err != nil {
		log.Println("Set failed", err)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	// add gameinfo
	player.JoinNewGameWithColor(gameInfo)
	invitePlayer.JoinNewGameWithColor(gameInfo)
	// save to db
	err = db.SetPlayerInfo(player.GetDbKey(), player)
	if err != nil {
		log.Println(err)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	err = db.SetPlayerInfo(player.GetDbKey(), invitePlayer)
	if err != nil {
		log.Println(err)
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	return &game.RESP_Weiqi_03{
		Status: conf.SUCCEED,
		GameId: gameInfo.WeiqiId,
	}
}
