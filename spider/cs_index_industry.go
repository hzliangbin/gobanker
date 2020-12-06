package spider

import (
	"fmt"
	"gobanker/models"
	"gobanker/util"
	"regexp"
	"strconv"
	"strings"
)

var url string = "http://www.csindex.com.cn/uploads/downloads/other/files/zh_CN/ZzhyflWz.zip"
var referrer string = "http://www.csindex.com.cn/zh-CN/downloads/industry-class"
var downloadFolder string = "./download/"



//TODO 后期更改成定时任务，每天执行一次
func CsIndexIndustryHandler() []models.CsIndexIndustry {
	path := util.DownloadFile(url,referrer, downloadFolder,"csindextype.zip")
	fmt.Println(path)
	err := util.DecompressZip(path, downloadFolder)
	if err != nil {
		fmt.Println(err)
	}
	res, err := util.ReadXLSData(downloadFolder +"cicslevel2.xls")
	if err != nil {
		fmt.Println(err)
	}
	dateTime, err := strconv.ParseInt(util.GetDateYYYYMMdd(),10,64)
	if err != nil {
		fmt.Println(err)
	}
	var data []models.CsIndexIndustry
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
		data = append(data, models.CsIndexIndustry{
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
	ctx, cli := util.ConnectWithColl("t_cs_index")
	defer func() {
		if err = cli.Close(ctx); err != nil {
			panic(err)
		}
	}()
	//此处全量删，再全量插入，股票有可能退市
	if err = cli.DropCollection(ctx); err != nil {
		fmt.Println("delete csindex collection error")
	}
	if _, err = cli.Collection.InsertMany(ctx, data); err != nil {
		fmt.Println("insert csindex collection error")
	}
	return  data
}