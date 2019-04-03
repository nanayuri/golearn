package main

import (
	"fmt"
	"os"

	"github.com/Luxurioust/excelize"
)

func main() {
	xlsx, err := excelize.OpenFile("C:/Users/Chao/Desktop/NBL3 NF IO CLASS BAS V2.8 (20190103).xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Get value from cell by given sheet index and axis
	cell, err := xlsx.GetCellValue("CLASS", "L3")
	fmt.Println(cell, err)
	// Get all the rows in a sheet.
	//index := xlsx.GetSheetIndex("CLASS")
	rows, _ := xlsx.GetRows("CLASS")

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
