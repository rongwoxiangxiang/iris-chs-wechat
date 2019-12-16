package script

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"strings"
)

func (script ScriptFuncs) FormatSqls(args []string) (ret interface{}) {
	file, err := os.OpenFile("account.txt", os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error[1]!", err)
		return
	}
	defer file.Close()

	if _, err := file.Stat(); err != nil {
		logrus.Fatalf("FormatLids err[2] %v", err)
	}
	i := 0
	buf := bufio.NewReader(file)
	for {
		i++
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
		buf := []byte(line)
		fileWrite, err := os.OpenFile("account/account"+strconv.Itoa(i)+".sql", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		writeObj := bufio.NewWriterSize(fileWrite, 4096)

		if _, err := writeObj.Write(buf); err == nil {
			fmt.Println("Successful: ", i)
		}
		if err := writeObj.Flush(); err != nil {
			logrus.Fatalf("FormatLids err[3] %v", err)
		}
		fileWrite.Close()
	}

	return
}
