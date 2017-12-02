package game

import (
	"time"
)

type WeiqiInfo struct {
	WeiqiId int64       //每场棋局独立Id
	Size    uint32      //分别为9*9 13*13 19*19
	StepLog [2][]uint32 //步数信息 0黑 1白//之后落子信息为X轴*19+y,以棋局左下角建立直角坐标系
	JoinLog []int64     //棋盘信息
	Player  [2]string   //玩家编号 index0 执白 index1 执黑
	IsEnd   bool        //是否完成
	Prunes  [][]uint32
	//index0 	对局双方0黑 1白
	//index1 	手数 len(steplog[0/1])+1
	//index2~ 	提子位置，计算方法同上
}

type PlayerInfo struct {
	PlayerId string //玩家ID
	AllWQ    map[int64][]uint32
	//key 棋盘ID value:index0:(0黑 1白)
	//index1 0未完成 1已经完成 并且胜利 2 完成并且失败 3 和棋
	CountDT map[uint32]time.Time
	//成就编号 成就获取时间
}

func NewOneGame(player []string, size int) *WeiqiInfo {
	var cc WeiqiInfo
	cc.Default()
	cc.Size = uint32(size)
	cc.JoinLog = make([]int64, size)
	cc.StepLog[0] = []uint32{}
	cc.StepLog[1] = []uint32{}
	newPlayer := [2]string{}
	for index := 0; index < 2; index++ {
		newPlayer[index] = player[index]
	}
	cc.Player = newPlayer
	return &cc
}
