package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func mqttGwDbinit() (err error) {
	dbAddrStr := fmt.Sprintf("%s:%d", config.Db.Host, config.Db.TcpPort)
	ue.db = redis.NewClient(&redis.Options{
		Addr:     dbAddrStr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return
}