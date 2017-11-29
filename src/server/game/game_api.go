package game

import (
	"conf"
	"fmt"
	"time"
)

func (g *WeiqiInfo) Default() {
	timeNow := time.Now().UnixNano()
	g.WeiqiId = uint32(timeNow)
	g.StepLog[0] = make([]uint32, 0, 19*19)
	g.StepLog[1] = make([]uint32, 0, 19*19)
	g.JoinLog = make([]int64, 0, 2)
	g.IsEnd = false
}

func (p *PlayerInfo) Default(pid string) {
	p.PlayerId = pid
	p.AllWQ = make(map[uint32][]uint32)
	p.CountDT = make(map[uint32]time.Time)
}

func (w *WeiqiInfo) SetStepLog(placeInfo uint32, playerId string) {
	_, isSucceed := w.GetWeiqiPlayerColor(playerId)
	if !isSucceed {
		fmt.Println("player is not right", playerId)
	}
	//TODO add weiqi-logic and place check

}

func (w *WeiqiInfo) GetWeiqiPlayerColor(playerId string) (uint32, bool) {
	if w.Player[0] == playerId {
		return conf.BLACK_PLAYER, true
	} else if w.Player[1] == playerId {
		return conf.WHITE_PLAYER, true
	}
	return 0, false
}
