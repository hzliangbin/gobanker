package spider

import (
	"fmt"
	"gobanker/util"
	"regexp"
	"strconv"
	"strings"
)

var url string = "http://www.csindex.com.cn/uploads/downloads/other/files/zh_CN/ZzhyflWz.zip"
var referrer string = "http://www.csindex.com.cn/zh-CN/downloads/industry-class"
var download_folder string = "./download/"

type CsIndexIndustry struct {
	Code string
	Name string
	LvOne string
	LvTwo string
	LvThree string
	LvFour string
	Date int64
}

func CsIndexIndustryHandler() []CsIndexIndustry {
	path := util.DownloadFile(url,referrer,download_folder,"csindextype.zip")
	fmt.Println(path)
	err := util.DecompressZip(path,download_folder)
	if err != nil {
		fmt.Println(err)
	}
	res, err := util.ReadXLSData(download_folder+"cicslevel2.xls")
	if err != nil {
		fmt.Println(err)
	}
	dateTime, err := strconv.ParseInt(util.GetDateYYYYMMdd(),10,64)
	if err != nil {
		fmt.Println(err)
	}
	var data []CsIndexIndustry
	r := regexp.MustCompile("\\s")
	for i, cells := range res {
		if i == 0 {
			continue
		}
		code := cells[0]
		ok, _ := regexp.MatchString("\\d{6}",code)
		//代码不为6位数字，或是900开头（b股）
		if !ok || strings.HasPrefix(code,"900") {
			continue
		}
		data = append(data, CsIndexIndustry{
			Code:    code,
			Name:    r.ReplaceAllString(cells[1],""),
			LvOne:   cells[5],
			LvTwo:   cells[11],
			LvThree: cells[14],
			LvFour:  cells[16],
			Date:    dateTime,
		})
	}
	for _, v := range data {
		fmt.Printf("%+v\r\n", v)
	}
	return  data
}