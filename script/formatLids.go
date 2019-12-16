package script

import (
	"bufio"
	"chs/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

func (script ScriptFuncs) FormatLids(args []string) (ret interface{}) {
	file, err := os.OpenFile(args[0], os.O_RDWR, 0666)
	fileWrite, err := os.OpenFile("userrwrite.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open file error[1]!", err)
		return
	}
	defer file.Close()

	if _, err := file.Stat(); err != nil {
		logrus.Fatalf("FormatLids err[2] %v", err)
	}

	defer fileWrite.Close()
	writeObj := bufio.NewWriterSize(fileWrite, 4096)

	l := 0
	str := ""
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		line = strings.TrimSpace(line)
		index := strings.IndexAny(line, "123456789")
		str += "'" + util.Substr(line, index, -2) + "',"

		l++
		if l%10000 == 0 {
			str = "SELECT `store_account` FROM `g_13002010`.`user_store` WHERE lid in(" + str[0:len(str)-1] + ");\n"
			buf := []byte(str)
			if _, err := writeObj.Write(buf); err == nil {
				fmt.Println("Successful: ", str)
			}
			str = ""
		}
	}
	if str != "" {
		str = "SELECT `store_account` FROM `g_13002010`.`user_store` WHERE lid in(" + str + ");\n"
		buf := []byte(str)
		if _, err := writeObj.Write(buf); err == nil {
			fmt.Println("Successful: ", str)
		}
		str = ""
	}
	if err := writeObj.Flush(); err != nil {
		logrus.Fatalf("FormatLids err[3] %v", err)
	}

	return
}
