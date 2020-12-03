package util

import (
	"github.com/shakinm/xlsReader/xls"
)
//TODO: make it more elegant and effective
func ReadXLSData(url string) ([][]string, error) {
	wb, err := xls.OpenFile(url)
	if err != nil {
		return nil, err
	}

	var res [][]string
	sheet, err := wb.GetSheet(0)
	if err != nil {
		return nil,err
	}
	for i := 0; i <= sheet.GetNumberRows(); i++ {
		if row, err := sheet.GetRow(i); err == nil {
			colsSize := len(row.GetCols())
			res = append(res,[]string{})
			for j := 0; j <= colsSize; j++ {
				if cell, err := row.GetCol(j); err == nil {
					res[i] = append(res[i], cell.GetString())
				}
			}
		}
	}
	return  res, nil
}

