package bin

import (
	"fmt"
	"server/cache"
	"time"
)

func LoadServerConf() {
	//init dfl
	cache.InitDflOnlineList()
}

func ClearOfflinePlayer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: %v", r)
		}
	}()
	//定时器 每10秒删除一遍在线玩家
	ticker := time.NewTicker(time.Second * 10)
	onlineList := cache.GetPlayerList()
	for {
		select {
		case <-ticker.C:
			go onlineList.RMOfflinePlayer()
		}
	}
}
