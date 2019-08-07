package script

import (
	"chs/config"
	"chs/modules/elasticsearch"
	"chs/util"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"reflect"
	"strconv"
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

func (script ScriptFuncs) BashDecryptPhones(args []string) (ret interface{}) {
	f, err := excelize.OpenFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	sheets := f.GetSheetMap()
	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err == nil {
			col, _ := excelize.ColumnNumberToName(len(rows[0]) + 2)
			for index, row := range rows {
				phone, _ := util.GetAesCryptor().Decrypt(row[1])
				f.SetCellStr(sheet, col+strconv.Itoa(index+1), phone)
			}
		}
	}
	f.Save()
	return
}

func InitEs(args []string) {
	config.InitConfig()
	elasticsearch.InitEs()
}
