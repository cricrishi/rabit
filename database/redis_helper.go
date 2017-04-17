
package database

import (
		"github.com/go-redis/redis"
		conf "github.com/rabit/config"
	
	)


func GetClient() (*redis.Client){
	c := conf.ConfigData["config"].Get("redis")
	redisConfig, _ := c.(map[string]interface{})
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig["host"].(string)+":" + redisConfig["port"].(string),
		Password: redisConfig["password"].(string), // no password set
		DB:       0,  // use default DB
	})
	return client
}


func SetKey(key string,value string,ttl int) error{
	
	redisClient := GetClient()
	
	err := redisClient.Set(key, value, 0).Err()
	return err
}

