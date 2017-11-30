package cache

import (
	"sync"
)

// 进行中的游戏结构体

type AllGameInfo struct {
	GameInfo map[int64]*SimpleGameLog
	//单纯的记录步数
	lock sync.Locker
}

type SimpleGameLog struct {
	GameLog []uint32
	IsEnd   bool
	lock    sync.Locker
}

func (sg *SimpleGameLog) Init() {
	if sg.GameLog == nil {
		sg.GameLog = make([]uint32, 0, 50)
	}
}

func (sg *SimpleGameLog) SetOneStep(stepNum uint32) {
	defer sg.lock.Unlock()
	sg.lock.Lock()
	sg.GameLog = append(sg.GameLog, stepNum)
}

func (sg *SimpleGameLog) EndGame() {
	//日志记录到数据库并且保存 ，在map中删除
	sg.IsEnd = true
}
