package util

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Reader struct {
	io.Reader
	Total int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Println("\r进度 %.2f%%", float64(r.Current * 10000/ r.Total)/100)
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

	if err = os.Mkdir(dir,0666); err != nil {
		fmt.Println("dir exists already")
	}

	f, err := os.Create(dir+filename)
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

func DecompressZip(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range  reader.File {
		rc, err := file.Open()
		if err != nil {
			return  err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil  {
			return err
		}
		w, err := os.Create(filename)
		if  err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w,rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path,"/"))
}

func subString(str string, start,end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}
	if end < start || end > length {
		panic("end is wrong")
	}
	return string(rs[start:end])
}