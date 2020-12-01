package util

import (
	"io"
	"net/http"
	"os"
	"time"
)
import "github.com/golang/glog"

type Reader struct {
	io.Reader
	Total int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	glog.Info("\r进度 %.2f%%", float64(r.Current * 10000/ r.Total)/100)
	return
}

func DownloadFile(url, referrer, dir, filename string) string {
	//r, err := http.Get(url)
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	r, err := http.NewRequest("GET",url,nil)
	if err != nil {
		panic(err)
	}
	r.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3346.9 Safari/537.36")
	r.Header.Set("Referer", referrer)

	res, err := client.Do(r)
	defer func() { _ = res.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close()}()

	reader := &Reader{
		Reader: res.Body,
		Total: res.ContentLength,
	}
	_, _ = io.Copy(f, reader)
	return dir+filename
}