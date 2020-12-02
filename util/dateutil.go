package util

import (
	"github.com/hzliangbin/date"
)



func GetDateYYYYMMdd() string {
	d := date.Now()
	res, err := d.Format("yyyyMMdd",false);
	if err != nil {
		return ""
	}
	return res
}