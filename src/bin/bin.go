package bin

import (
	"fmt"
	sbin "server/bin"
	"server/cache"
	"time"
)

func LoadServerConf() {
	//init dfl
	cache.InitDflOnlineList()
	cache.InitDflMatchList()
}

func ClearOfflinePlayer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: %v", r)
		}
	}()
	//定时器 每10秒删除一遍在线玩家
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			go cache.RMOfflinePlayerForTick()
		}
	}
}

func MatchPlayer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic: %v", r)
		}
	}()
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			go MatchPlayerBySize()
		}
	}
}

func MatchPlayerBySize() {
	gameList := cache.MatchGameBySize9()
	for _, v := range gameList {
		sbin.Weiqi03(v[0], v[1], 9)
	}
	gameList = cache.MatchGameBySize13()
	for _, v := range gameList {
		sbin.Weiqi03(v[0], v[1], 13)
	}
	gameList = cache.MatchGameBySize19()
	for _, v := range gameList {
		sbin.Weiqi03(v[0], v[1], 19)
	}
}
