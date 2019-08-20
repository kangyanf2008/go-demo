package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Stime2TimeStamp(FormatTime(time.Now())))
}


func FormatTime(t time.Time) string {
	f := "2006-01-02 15:04:05"
	return t.Format(f)
}

//字符串转换为时间戳
func Stime2TimeStamp(stime string) int64 {
	f := "2006-01-02 15:04:05"
	t, e := time.ParseInLocation(f, stime, time.Local)
	if e != nil {
		//logging.Logger.Errorf("string to time stamp stime=[%s],error=[%s]", stime, e)
	}
	return t.UnixNano() / 1e6
}

//获取当前时间戳，为毫秒
func TimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}
