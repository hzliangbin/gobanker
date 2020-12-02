package util

import (
	"github.com/extrame/xls"
)

func ReadXLSData(url string) ([][]string, error) {
	xlsFile, err := xls.Open(url,"utf-8")
	if err != nil {
		return nil, err
	}
	numOfSheets := xlsFile.NumSheets()
	var res [][]string

	for i := 0; i < numOfSheets; i++ {
		if sheet := xlsFile.GetSheet(i); sheet != nil {
			rowNum := int(sheet.MaxRow)
			for j := 0; j <= rowNum; j++ {
				row := sheet.Row(j)
				colNum := row.LastCol()
				ans := make([]string, colNum)
				for k := 0; k < colNum; k++ {
					ans[i] = row.Col(k)
				}
				res = append(res, ans)
			}
		}
	}
	return  res, nil
}