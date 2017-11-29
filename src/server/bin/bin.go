package bin

import (
	"conf"
	"fmt"
	"server/cache"
	"server/db"
)

func Weiqi01(playerId string) (uint32, []string) {
	player, err := db.GetPlayerInfo(playerId)
	if err != nil {
		fmt.Println("The first time to login in", playerId)
		player.Default(playerId)
	}
	//add into PlayerList
	err = db.SetAllPlayerIdList(playerId)
	if err != nil {
		fmt.Println(err)
		return conf.ERR_SERVER_ERR, nil
	}
	cache.OutAddOnlinePlayer(player.PlayerId)
	onlineList := cache.GetAllOnlinePlayer(playerId)
	return conf.SUCCEED, onlineList
}
