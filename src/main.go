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
	go bin.MatchPlayer()
	runGameServer()
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
		c.JSON(200, gin.H{
			"status":       respInfo.Status,
			"onlinePlayer": respInfo.OnlinePlayer,
			"allGameInfo":  respInfo.AllGameInfo,
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
		c.JSON(200, gin.H{
			"status": respInfo.Status,
			"gameid": respInfo.GameId,
		})
	})

	r.POST("/Weiqi04", func(c *gin.Context) {
		playerId := c.PostForm("Uid")
		gameId := c.PostForm("GameId")
		nextStepStr := c.PostForm("NextStep")
		nextStep, _ := strconv.Atoi(nextStepStr)
		respInfo := sbin.Weiqi04(playerId, gameId, nextStep)
		log.Println("/Weiqi04_RESP_INFO:", respInfo)
		c.JSON(200, gin.H{
			"status":     respInfo.Status,
			"gamestatis": respInfo.GameStatus,
		})
	})

	r.POST("/Weiqi05", func(c *gin.Context) {
		playerId := c.PostForm("Uid")
		gameId := c.PostForm("GameId")
		typeId := c.PostForm("Type")
		statusCode := sbin.Weiqi05(playerId, gameId, typeId)
		log.Println("/Weiqi05_RESP_INFO:", statusCode)
		c.JSON(200, gin.H{
			"status": statusCode,
		})
	})

	r.POST("/Weiqi06", func(c *gin.Context) {
		playerId := c.PostForm("Uid")
		gameId := c.PostForm("GameId")
		respInfo := sbin.Weiqi06(playerId, gameId)
		log.Println("/Weiqi06_RESP_INFO:", respInfo)
		c.JSON(200, gin.H{
			"status":     respInfo.Status,
			"size":       respInfo.Size,
			"playInfo":   respInfo.Player,
			"round":      respInfo.Round,
			"gameStatus": respInfo.GameStatus,
		})
	})
	r.POST("/Weiqi07", func(c *gin.Context) {
		playerId := c.PostForm("Uid")
		matchType := c.PostForm("Type")
		gameSize := c.PostForm("Size")
		statusCode := sbin.Weiqi07(playerId, matchType, gameSize)
		log.Println("/Weiqi07_RESP_INFO:", statusCode)
		c.JSON(200, gin.H{
			"status": statusCode,
		})
	})
	r.Run(":10087")
}
