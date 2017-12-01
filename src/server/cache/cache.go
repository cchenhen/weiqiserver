package cache

import (
	"sync"
	"time"
)

var dflOnlineList OnlineList

// 对该字段进行操作需要加锁
type OnlineList struct {
	List map[string]int64
	lock sync.RWMutex
}

func InitDflOnlineList() {
	dflOnlineList.List = make(map[string]int64)
}

func RMOfflinePlayerForTick() {
	onlineList := getPlayerList()
	onlineList.RMOfflinePlayer()
}

func getPlayerList() *OnlineList {
	return &dflOnlineList
}

// 加入活动
func (ol *OnlineList) AddOnlinePlayer(playerId string) {
	defer ol.lock.Unlock()
	ol.lock.Lock()
	score := time.Now().Unix() + 5
	dflOnlineList.List[playerId] = score
}

func OutAddOnlinePlayer(playerId string) {
	dflOnlineList.AddOnlinePlayer(playerId)
}

// 删除离线玩家 需要定时器每一秒钟执行一次
func (ol *OnlineList) RMOfflinePlayer() {
	defer ol.lock.Unlock()
	ol.lock.Lock()
	timeNow := time.Now().Unix()
	for k, v := range dflOnlineList.List {
		if timeNow > v {
			delete(dflOnlineList.List, k)
		}
	}
}

// 检查玩家是否在线
func (ol *OnlineList) CheckPlayerIsOnline(playerId string) bool {
	_, ok := ol.List[playerId]
	return ok
}

// 获取所有在线好友 除开自己
func GetAllOnlinePlayer(playerId string) []string {
	dflOnlineList.RMOfflinePlayer()
	onlineList := []string{}
	for k := range dflOnlineList.List {
		if playerId != k {
			onlineList = append(onlineList, playerId)
		}
	}
	return onlineList
}

func IsPlayerOnline(playerId string) bool {
	return dflOnlineList.CheckPlayerIsOnline(playerId)
}
