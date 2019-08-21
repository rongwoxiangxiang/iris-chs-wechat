package main

import (
	"chs/bootstrap"
	"github.com/go-redis/redis"
	"log"
)

func main() {
	log.Println("Init Application chs")
	bootstrap.App().Run()
}
