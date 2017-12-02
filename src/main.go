package main

import (
	"bin"
	"fmt"
	"log"
	sbin "server/bin"
	_ "server/cache"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	bin.LoadServerConf()
	//go bin.ClearOfflinePlayer()
	go runGameServer()
	for {

	}
}

func runGameServer() {
	r := gin.Default()
	//test
	r.POST("/Weiqi", func(c *gin.Context) {
		//解析POST中的内容
		id := c.PostForm("Uid")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": id,
		})
	})
	r.POST("/Weiqi01", func(c *gin.Context) {
		// encode data
		playerId := c.PostForm("Uid")
		respInfo := sbin.Weiqi01(playerId)
		log.Println("/Weiqi01_RESP_INFO:", respInfo)
		statusCodeStr := fmt.Sprintln(respInfo.Status)
		onlinelistStr := fmt.Sprintln(respInfo.OnlinePlayer)
		// return data
		c.JSON(200, gin.H{
			"status":       statusCodeStr,
			"onlinePlayer": onlinelistStr,
		})
	})
	r.POST("/Weiqi02", func(c *gin.Context) {
		// encode data
		playerId := c.PostForm("Uid")
		respInfo := sbin.Weiqi02(playerId)
		log.Println("/Weiqi02_RESP_INFO:", respInfo)
		statusCodeStr := fmt.Sprintln(respInfo.Status)
		LiveGame := fmt.Sprintln(respInfo.LiveGame)
		onlinePlayer := fmt.Sprintln(respInfo.OnlinePlayer)
		// return data
		c.JSON(200, gin.H{
			"status":       statusCodeStr,
			"liveGame":     LiveGame,
			"onlinePlayer": onlinePlayer,
		})
	})
	r.POST("/Weiqi03", func(c *gin.Context) {
		// encode data
		playerId := c.PostForm("Uid")
		inviteId := c.PostForm("InviteId")
		sizeStr := c.PostForm("Size")
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			log.Println("BadParam", sizeStr)
			size = 0
		}
		respInfo := sbin.Weiqi03(playerId, inviteId, size)
		log.Println("/Weiqi03_RESP_INFO:", respInfo)
		statusCodeStr := fmt.Sprintln(respInfo.Status)
		gameId := fmt.Sprintln(respInfo.GameId)
		c.JSON(200, gin.H{
			"status": statusCodeStr,
			"gameid": gameId,
		})
	})

	r.POST("/Weiqi04", func(c *gin.Context) {
		playerId := c.PostForm("Uid")
		gameId := c.PostForm("GameId")
		nextStepStr := c.PostForm("NextStep")
		nextStep, _ := strconv.Atoi(nextStepStr)
		respInfo := sbin.Weiqi04(playerId, gameId, nextStep)
		log.Println("/Weiqi04_RESP_INFO:", respInfo)
		statusCodeStr := fmt.Sprintln(respInfo.Status)
		gameInfo := fmt.Sprintln(respInfo.GameStatus)
		c.JSON(200, gin.H{
			"status":     statusCodeStr,
			"gamestatis": gameInfo,
		})
	})

	r.Run(":10087")
}
