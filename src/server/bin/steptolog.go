package bin

func StepToGameInfo(gameStep []int64) [][]uint32 {
	gameSize := len(gameStep)
	gameInfo := make([][]uint32, gameSize)
	for k := range gameInfo {
		gameInfo[k] = make([]uint32, gameSize)
	}
	for i := 0; i < gameSize; i++ {
		row := gameStep[i]
		for j := uint(0); j < uint(gameSize*2); j += 2 {
			a := row
			b := a << (61 - j)
			c := b >> (61)
			switch c {
			case 0:
				//空白子
				gameInfo[i][j] = 0
			case 2:
				//黑子
				gameInfo[i][j] = 1
			case 3:
				//白子
				gameInfo[i][j] = 2
			}
		}
	}
	return gameInfo
}

func StepLogToGameShow(gameStep [][]uint32) []int64 {
	sizeLen := len(gameStep)
	newJoinLog := make([]int64, sizeLen)
	for i := 0; i < sizeLen; i++ {
		newLog := int64(0)
		for j := uint(0); j < uint(sizeLen*2); j += 2 {
			switch gameStep[i][j] {
			case 0:
				newLog |= 0 << j * 2
				newLog |= 0<<j*2 + 1
			case 2:
				newLog |= 1 << j * 2
				newLog |= 0<<j*2 + 1
			case 3:
				newLog |= 1 << j * 2
				newLog |= 1<<j*2 + 1
			}
		}
		newJoinLog[i] = newLog
	}
	return newJoinLog
}
