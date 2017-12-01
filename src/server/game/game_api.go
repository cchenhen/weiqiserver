package game

import (
	"conf"
	"fmt"
	"time"
)

func (g *WeiqiInfo) Default() {
	g.WeiqiId = time.Now().UnixNano()
	g.StepLog[0] = make([]uint32, 0, 19*19)
	g.StepLog[1] = make([]uint32, 0, 19*19)
	g.JoinLog = make([]int64, 0, 2)
	g.IsEnd = false
}

func (g *WeiqiInfo) GetDbKey() string {
	return fmt.Sprintf("Weiqi:Game:%v", g.WeiqiId)
}

func (p *PlayerInfo) Default(pid string) {
	p.PlayerId = pid
	p.AllWQ = make(map[int64][]uint32)
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

// 生成Redis的key
func (p *PlayerInfo) GetDbKey() string {
	return fmt.Sprintf("Weiqi:Player:%v", p.PlayerId)
}

// 加入新生成的棋局
func (p *PlayerInfo) JoinNewGameWithColor(gameInfo *WeiqiInfo) {
	color, _ := gameInfo.GetWeiqiPlayerColor(p.PlayerId)
	p.AllWQ[gameInfo.WeiqiId] = []uint32{color, conf.UN_FINISHED}
}

// GetOnGame 获取进行中的棋局
func (p *PlayerInfo) GetOnGame() []int64 {
	if p.AllWQ == nil {
		return nil
	}
	allGame := make([]int64, 0, len(p.AllWQ))
	for k, v := range p.AllWQ {
		if v[1] == conf.UN_FINISHED {
			allGame = append(allGame, k)
		}
	}
	return allGame
}
