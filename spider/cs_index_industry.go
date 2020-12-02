package spider

import (
	"fmt"
	"gobanker/util"
	"strconv"
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

func CsIndexIndustryHandler() {
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
	fmt.Println(res[0][0])
	fmt.Println(dateTime)



}