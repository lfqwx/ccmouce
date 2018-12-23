package read

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Read()[][]string  {
	xlsx, err := excelize.OpenFile("./cycle.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	// Get value from cell by given worksheet name and axis.
	rows := xlsx.GetRows("Sheet1")
	return rows
}
