package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// 获取工作表中指定单元格的值
	cell := f.GetCellValue("BigThings", "D2")
	println(cell)
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows("BigThings")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}

	f.SetCellValue("BigThings", "H1", "T7 Hello world.")
	f.SetCellValue("BigThings", "H2", 666)
	f.UpdateLinkedValue()
	err = f.Save()
	if err != nil {
		print(err)
	}
}
