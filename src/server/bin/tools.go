package bin

import (
	"fmt"
	"log"
	"server/db"
	"server/game"
)

func GetColorByPlayId(playerInfo *game.PlayerInfo, gameId int64) uint32 {
	key := fmt.Sprintf("Weiqi:Game:%v", gameId)
	gameInfo, err := db.GetRedisC(key)
	if err != nil {
		log.Println("Bad GameId", gameId)
		return 0
	}
	return gameInfo.GetNextStepColor()
}

func GetAllOnlineGameInfo(liveGame []int64) []game.AllGameInfo {
	allGameInfo := make([]game.AllGameInfo, 0, len(liveGame))
	for _, v := range liveGame {
		key := fmt.Sprintf("Weiqi:Game:%v", v)
		gameInfo, err := db.GetRedisC(key)
		if err != nil {
			log.Println("Bad GameId", key)
			return nil
		}
		newGameInfo := &game.AllGameInfo{}
		newGameInfo.GameId = fmt.Sprintln(v)
		newGameInfo.Size = gameInfo.Size
		newGameInfo.Round = gameInfo.GetNextStepColor()
		newGameInfo.Player = gameInfo.Player
		newGameInfo.GameStatus = gameInfo.JoinLog
		allGameInfo = append(allGameInfo, *newGameInfo)
	}
	return allGameInfo
}
