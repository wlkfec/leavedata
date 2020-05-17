package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ReadExcel() {
	// 打开文件
	xlFile, err := xlsx.OpenFile("old.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历sheet页读取
	for _, sheet := range xlFile.Sheets {
		fmt.Println("sheet name: ", sheet.Name)
		//遍历行读取
		for _, row := range sheet.Rows {
			fmt.Println(row.Cells[3].String())
			// 遍历每行的列读取
			//for _, cell := range row.Cells {
			//	text := cell.String()
			//	fmt.Printf("%20s", text)
			//}
			//fmt.Print("\n")
		}
	}
}
