package main

import (
	"bin"
	"fmt"
	"log"
	sbin "server/bin"
	_ "server/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to redis test
	// load server conf
	// add log conf to show filename and fileline
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	bin.LoadServerConf()
	//
	go bin.ClearOfflinePlayer()
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
		respInfo := sbin.Weiqi03(playerId, inviteId)
		statusCodeStr := fmt.Sprintln(respInfo.Status)
		gameId := fmt.Sprintln(respInfo.GameId)
		c.JSON(200, gin.H{
			"status": statusCodeStr,
			"gameid": gameId,
		})
	})
	r.Run(":10087")
}
