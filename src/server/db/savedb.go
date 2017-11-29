package db

import (
	"encoding/json"
	"fmt"
	"server/game"

	"github.com/garyburd/redigo/redis"
)

// 写入
func SetRedisC(key string, userinfo *game.WeiqiInfo) error {
	//TODO add read file
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}

	defer c.Close()

	// 写入
	_, err = c.Do("SET", key, userinfo)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	return err
}

// 取出游戏数据
func GetRedisC(key string) (*game.WeiqiInfo, error) {
	//TODO add read file
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}

	defer c.Close()
	// 获取
	key = fmt.Sprintf("weiqi:game:%v", key)
	weiqiInfo, err := c.Do("GET", key)
	if err != nil {
		fmt.Println("redis get failed:", err)
		return nil, err
	}
	var wi game.WeiqiInfo
	weiqiStr := fmt.Sprintf("%v", weiqiInfo)
	err = json.Unmarshal([]byte(weiqiStr), &wi)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return nil, err
	}
	return &wi, err
}

// 插入所有游戏用户列表
func SetAllPlayerIdList(playerId string) error {
	//TODO add read file
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return err
	}

	defer c.Close()

	key := "weiqi_server_allplayer"
	_, err = c.Do("SADD", key, playerId)
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	return nil
}
