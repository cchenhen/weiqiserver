package bin

import (
	"center"
	"conf"
	"fmt"
	"log"
	"server/cache"
	"server/db"
	"server/game"
)

func Weiqi01(playerId string) *game.RESP_Weiqi_01 {
	player, err := db.GetPlayerInfo(playerId)
	if err != nil {
		log.Println("The first time to login in:", playerId)
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

func Weiqi03(playerId string, inviteId string, size int) *game.RESP_Weiqi_03 {
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
	// check size
	if size == conf.WEIQI_SIZE_SMALL || size == conf.WEIQI_SIZE_MID || size == conf.WEIQI_SIZE_STANDARD {
		return &game.RESP_Weiqi_03{
			Status: conf.ERR_BAD_PARAM,
		}
	}
	// make a new game
	playerList := []string{playerId, inviteId}
	gameInfo := game.NewOneGame(playerList, size)
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

func Weiqi04(playerId string, gameId string, nextStep int) *game.RESP_Weiqi_04 {
	_, err := db.GetPlayerInfo(playerId)
	if err != nil {
		log.Println("Bad PlayerId", playerId)
		return &game.RESP_Weiqi_04{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	key := fmt.Sprintf("Weiqi:Game:%v", gameId)
	gameInfo, err := db.GetRedisC(key)
	if err != nil {
		log.Println("Bad GameId", gameId)
		return &game.RESP_Weiqi_04{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	//check is the right step
	nextStepColor := gameInfo.GetNextStepColor()
	playerColor, _ := gameInfo.GetWeiqiPlayerColor(playerId)
	if nextStepColor != playerColor {
		log.Println("Bad Game Step")
		return &game.RESP_Weiqi_04{
			Status: conf.ERR_BAD_PARAM,
		}
	}
	gameInfo.AddOneLogStep(nextStepColor, nextStep)
	x := nextStep / 19
	y := nextStep % 19
	// JoinLog change to [size][size]uint32
	gameLogStep := StepToGameInfo(gameInfo.JoinLog)
	if nextStep != conf.GIVE_UP {
		gameLogStep[x][y] = nextStepColor + 1
	}
	// 进行提子
	newGameLogStep := center.GameCenterLogic(gameLogStep, nextStepColor, gameInfo.Size)
	newJoinStep := StepLogToGameShow(newGameLogStep)
	gameInfo.JoinLog = newJoinStep
	//save db
	// save to db
	err = db.SetRedisC(gameInfo.GetDbKey(), gameInfo)
	if err != nil {
		log.Println("Set failed", err)
		return &game.RESP_Weiqi_04{
			Status: conf.ERR_SERVER_ERR,
		}
	}
	return &game.RESP_Weiqi_04{
		Status:     conf.ERR_SERVER_ERR,
		GameStatus: newJoinStep,
	}
}
