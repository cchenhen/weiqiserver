package center

import "conf"

// 围棋提子逻辑

func GameCenterLogic(gameInfo [][]uint32, getColor uint32, size uint32) [][]uint32 {
	TakeDeadChess(gameInfo, getColor, size)
	return gameInfo
}

func TakeDeadChess(gameInfo [][]uint32, getColor uint32, gameSize uint32) {
	allInfo := make([][2]uint32, gameSize*gameSize)
	hasDealWith := make([][]bool, gameSize)
	for k := range hasDealWith {
		hasDealWith[k] = make([]bool, gameSize)
	}
	var (
		top int
		pos int
	)
	for i := uint(0); i < uint(gameSize); i++ {
		for j := uint(0); j < uint(gameSize); j++ {
			if gameInfo[i][j] != getColor || hasDealWith[i][j] {
				continue
			}
			hasDealWith[i][j] = true
			allInfo[top][0] = uint32(i)
			allInfo[top][1] = uint32(j)
			top++
			for pos < top {
				x := allInfo[pos][0]
				y := allInfo[pos][1]
				if x > 0 && !hasDealWith[x-1][y] && getColor == gameInfo[x-1][y] {
					hasDealWith[x-1][y] = true
					allInfo[top][0] = x - 1
					allInfo[top][1] = y
					top++
				}
				if x < (gameSize-1) && !hasDealWith[x+1][y] && getColor == gameInfo[x+1][y] {
					hasDealWith[x+1][y] = true
					allInfo[top][0] = x + 1
					allInfo[top][1] = y
					top++
				}
				if y > 0 && !hasDealWith[x][y-1] && getColor == gameInfo[x][y-1] {
					hasDealWith[x][y-1] = true
					allInfo[top][0] = x
					allInfo[top][1] = y - 1
					top++
				}
				if y < (gameSize-1) && !hasDealWith[x][y+1] && getColor == gameInfo[x][y+1] {
					hasDealWith[x][y+1] = true
					allInfo[top][0] = x
					allInfo[top][1] = y + 1
					top++
				}
				pos++ //next step
			}
			//TODO is dead chess
			if top > 0 && IsDeadChess(gameInfo, allInfo, top, gameSize) {
				ClearDeadChess(gameInfo, allInfo, top)
			}
			top = 0
			pos = 0
		}
	}
}

func IsDeadChess(gameInfo [][]uint32, allInfo [][2]uint32, top int, gameSize uint32) bool {
	var (
		x, y uint32
	)
	for i := 0; i < top; i++ {
		x = allInfo[i][0]
		y = allInfo[i][1]
		if x > 0 && gameInfo[x-1][y] == conf.SPACE_G {
			return false
		}
		if x < gameSize-1 && gameInfo[x+1][y] == conf.SPACE_G {
			return false
		}
		if y > 0 && gameInfo[x][y-1] == conf.SPACE_G {
			return false
		}
		if y < gameSize-1 && gameInfo[x][y+1] == conf.SPACE_G {
			return false
		}
	}
	return true
}

// 清除死子
func ClearDeadChess(gameInfo [][]uint32, allInfo [][2]uint32, top int) {
	for i := 0; i < top; i++ {
		gameInfo[allInfo[i][0]][allInfo[i][1]] = conf.SPACE_G
	}
}
