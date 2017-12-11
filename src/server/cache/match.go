package cache

import (
	"conf"
	"sync"
)

var dflMatchListSize9 MatchList
var dflMatchListSize13 MatchList
var dflMatchListSize19 MatchList

type MatchList struct {
	GameList map[string]uint // value 代表合法用户数量
	lock     sync.RWMutex
}

func InitDflMatchList() {
	if dflMatchListSize9.GameList == nil {
		dflMatchListSize9.GameList = make(map[string]uint)
	}
	if dflMatchListSize13.GameList == nil {
		dflMatchListSize13.GameList = make(map[string]uint)
	}
	if dflMatchListSize19.GameList == nil {
		dflMatchListSize19.GameList = make(map[string]uint)
	}
}

func (ml *MatchList) AddOnePlayerIntoMatchList(playerId string) {
	defer ml.lock.Unlock()
	ml.lock.Lock()
	ml.GameList[playerId] = 0
}

func (ml *MatchList) DeletePlayer(playerId string) {
	defer ml.lock.Unlock()
	ml.lock.Lock()
	delete(ml.GameList, playerId)
}

func (ml *MatchList) MatchPlayer(gameSize int) [][]string {
	defer ml.lock.Unlock()
	ml.lock.Lock()
	var playerA, playerB string
	matchResult := [][]string{}
	for k := range ml.GameList {
		if playerA == "" && playerB == "" {
			playerA = k
		} else if playerA != "" && playerB == "" {
			playerB = k
			// matching succeed
			matchResult = append(matchResult, []string{playerA, playerB})
			//clear playerA and playerB
			delete(ml.GameList, playerA)
			delete(ml.GameList, playerB)
			playerA = ""
			playerB = ""
		}
	}
	return matchResult
}

func (ml *MatchList) IsEndMatch(playId string) bool {
	if _, ok := ml.GameList[playId]; ok {
		return false
	}
	return true
}

func MatchGameBySize9() [][]string {
	return dflMatchListSize9.MatchPlayer(9)
}

func MatchGameBySize13() [][]string {
	return dflMatchListSize13.MatchPlayer(13)
}

func MatchGameBySize19() [][]string {
	return dflMatchListSize19.MatchPlayer(19)
}

func AddOnePlayerBySize(playerId string, gameSize int) {
	switch gameSize {
	case conf.WEIQI_SIZE_SMALL:
		dflMatchListSize9.AddOnePlayerIntoMatchList(playerId)
	case conf.WEIQI_SIZE_MID:
		dflMatchListSize13.AddOnePlayerIntoMatchList(playerId)
	case conf.WEIQI_SIZE_STANDARD:
		dflMatchListSize19.AddOnePlayerIntoMatchList(playerId)
	}
}

func EndMatchBySize(playerId string, gameSize int) {
	switch gameSize {
	case conf.WEIQI_SIZE_SMALL:
		dflMatchListSize9.DeletePlayer(playerId)
	case conf.WEIQI_SIZE_MID:
		dflMatchListSize13.DeletePlayer(playerId)
	case conf.WEIQI_SIZE_STANDARD:
		dflMatchListSize19.DeletePlayer(playerId)
	}
}

func GetMatchStatusByPlayerId(playerId string) uint32 {
	a := dflMatchListSize9.IsEndMatch(playerId)
	b := dflMatchListSize13.IsEndMatch(playerId)
	c := dflMatchListSize19.IsEndMatch(playerId)
	if a && b && c {
		return 0
	}
	return 1
}
