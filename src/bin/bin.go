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
	ticker := time.NewTicker(time.Second * 1)
	onlineList := cache.GetPlayerList()
	for {
		select {
		case <-ticker.C:
			go onlineList.RMOfflinePlayer()
		}
	}
}
