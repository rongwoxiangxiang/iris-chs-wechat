package script

import (
	"chs/util"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

func Decrypt() {
	f, err := excelize.OpenFile("./phonesbak.xlsx")
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
}

func getAlias() {

}
