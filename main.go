package main

import (
	"chs/bootstrap"
	"log"
)

func main() {
	log.Println("Init Application chs")
	bootstrap.App().Run()
}
