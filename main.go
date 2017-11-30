package main

import (
	"convertable/base"
	"os"

	"github.com/tealeg/xlsx"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("請輸入格式:convertable [*.xlsx]")
	// 	os.Exit(-1)
	// }

	// fileName := conf.Setting.FilePath + os.Args[1]
	fileName := "./error_code_table/error_code.xlsx"

	xlFile, err := xlsx.OpenFile(fileName)
	base.Checkerr(err)
	xlsxColumns := xlFile.Sheets[0].Rows[0].Cells
	rowsnum := len(xlFile.Sheets[0].Rows)
	topString := "//This file is creating automatic, don't edit!\n"
	constString := "package constant\n" +
		"const (\n"
	nameString := "//ErrorCodeName is a map to get error string by error code\n" + "var ErrorCodeName = map[int]string{\n"
	valueString := "//ErrorCodeValue is a map to get error code by error string\n" + "var ErrorCodeValue = map[string]int{\n"
	for i := 1; i < rowsnum; i++ {
		for k, _ := range xlsxColumns {
			if k == 0 {
				continue //跳過流水號
			}

		}
		errorName := xlFile.Sheets[0].Rows[i].Cells[1].String()
		errorCode := xlFile.Sheets[0].Rows[i].Cells[2].String()
		comment := xlFile.Sheets[0].Rows[i].Cells[3].String()
		constString += "    //" + comment + "\n    " + errorName + " int = " + errorCode + "\n"
		nameString += "    " + errorCode + ": \"" + errorName + "\",\n"
		valueString += "    \"" + errorName + "\":" + errorCode + ",\n"
	}

	constString += ")\n"
	nameString += "}\n"
	valueString += "}\n"

	f, _ := os.OpenFile("./constant/errorcode.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	f.WriteString(topString + constString + nameString + valueString)
}
