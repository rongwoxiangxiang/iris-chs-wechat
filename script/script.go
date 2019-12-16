package script

import (
	"bufio"
	_ "chs/config"
	"chs/modules/elasticsearch"
	"chs/util"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Do(method string, args []string) {
	scripts := ScriptFuncs{time.Now()}
	scriptsF := reflect.ValueOf(&scripts).Elem()
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(args)
	scriptsF.MethodByName(method).Call(params)
}

type ScriptFuncs struct {
	startTime time.Time
}

func (script ScriptFuncs) BashDecryptPhonesTxt(args []string) (ret interface{}) {
	file, err := os.OpenFile(args[0], os.O_RDWR, 0666)
	file2, err := os.OpenFile("res.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		phone, err2 := util.GetAesCryptor().Decrypt(line)
		if err2 != nil {
			phone = err2.Error()
		}
		file2.WriteString(phone + "\n")
		fmt.Println(line + "               " + phone)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	return
}

func (script ScriptFuncs) BashDecryptPhonesXls(args []string) (ret interface{}) {
	f, err := excelize.OpenFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	sheets := f.GetSheetMap()
	for _, sheet := range sheets {
		rows := f.GetRows(sheet)
		if err == nil {
			col, _ := columnNumberToName(len(rows[0]) + 2)
			for index, row := range rows {
				phone, _ := util.GetAesCryptor().Decrypt(row[1])
				f.SetCellStr(sheet, col+strconv.Itoa(index+1), phone)
			}
		}
	}
	f.Save()
	return
}

func columnNumberToName(num int) (string, error) {
	if num < 1 {
		return "", fmt.Errorf("incorrect column number %d", num)
	}
	var col string
	for num > 0 {
		col = string((num-1)%26+65) + col
		num = (num - 1) / 26
	}
	return col, nil
}

func InitEs(args []string) {
	elasticsearch.InitEs()
}
