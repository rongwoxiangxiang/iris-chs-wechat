package main

import (
	"bytes"
	"fmt"
	"os"
)

var (
	EXIST_COMMANDS_MAP = map[string]string{
		"decrypt": "BashDecryptPhones",
		"init-es": "InitEs",
	}
)

func main() {
	var b bytes.Buffer
	str := "11111"
	bt := []byte(str)
	b.Write(bt)
	fmt.Fprint(&b, "22222")
	b.WriteTo(os.Stdout)

	os.Exit(0)

	//args := os.Args
	//if len(args) < 2 {
	//	log.Fatal("Commands: args not be empty")
	//	return
	//}
	//method, ok := EXIST_COMMANDS_MAP[args[1]]
	//if !ok {
	//	log.Fatal("Commands: command not exist")
	//	return
	//}
	//script.Do(method, args[2:])
}
