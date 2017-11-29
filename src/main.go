package main

import (
	"bin"
	"fmt"
	sbin "server/bin"
	_ "server/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect to redis test
	// load server conf
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
		// TODO game logic
		statusCode, onlinelist := sbin.Weiqi01(playerId)
		statusCodeStr := fmt.Sprintln(statusCode)
		onlinelistStr := fmt.Sprintln(onlinelist)
		// return data
		c.JSON(200, gin.H{
			"status":       statusCodeStr,
			"onlinePlayer": onlinelistStr,
		})
	})
	r.Run(":10087")
}
