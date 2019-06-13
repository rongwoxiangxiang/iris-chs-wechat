package main

import (
	"chs/bootstrap"
	"go-common/library/log"
)

func main()  {
	log.Info("Init Application chs")
	bootstrap.App().Run()
}
