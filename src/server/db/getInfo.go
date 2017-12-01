package db

import (
	"encoding/json"
	"fmt"
	"log"
	"server/game"

	"github.com/garyburd/redigo/redis"
)

func GetPlayerInfo(key string) (*game.PlayerInfo, error) {
	//TODO add read file
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("Connect to redis error", err)
		return nil, err
	}

	defer c.Close()
	// 获取
	key = fmt.Sprintf("Weiqi:Player:%v", key)
	piInfo, err := c.Do("GET", key)
	if err != nil {
		log.Println("redis get failed:", err)
		return nil, err
	}
	var pi game.PlayerInfo
	b, err := redis.Bytes(piInfo, nil)
	if err != nil {
		log.Println("redis format failed:", err)
		return nil, err
	}
	if len(b) == 0 {
		log.Println("redis db miss")
		return nil, err
	}
	err = json.Unmarshal(b, &pi)
	if err != nil {
		log.Println("redis get failed:", err)
		return nil, err
	}
	return &pi, err
}

func SetPlayerInfo(key string, userinfo *game.PlayerInfo) error {
	//TODO add read file
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("Connect to redis error", err)
		return err
	}

	defer c.Close()

	// 写入
	jsonByte, err := json.Marshal(userinfo)
	_, err = c.Do("SET", key, jsonByte)
	if err != nil {
		log.Println("redis set failed:", err)
		return err
	}
	return err
}
