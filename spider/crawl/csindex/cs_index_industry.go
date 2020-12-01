package csindex

import (
	"github.com/golang/glog"
	"gobanker/util"
)

var url string = "http://www.csindex.com.cn/uploads/downloads/other/files/zh_CN/ZzhyflWz.zip"
var referrer string = "http://www.csindex.com.cn/zh-CN/downloads/industry-class"
var download_folder string = "./download/"

func Handler() {
	path := util.DownloadFile(url,referrer,download_folder,"csindextype.zip")
	glog.Info(path)
}