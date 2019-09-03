package main

import (
	"chs/script"
	"log"
	"os"
)

var (
	EXIST_COMMANDS_MAP = map[string]string{
		"decrypt": "BashDecryptPhones",
		"init-es": "InitEs",
	}
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Commands: args not be empty")
		return
	}
	method, ok := EXIST_COMMANDS_MAP[args[1]]
	if !ok {
		log.Fatal("Commands: command not exist")
		return
	}
	script.Do(method, args[2:])
}
