package main

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func main() {
/*	fmt.Println(Stime2TimeStamp(FormatTime(time.Now())))
	fmt.Println(FormatTime(time.Now().UTC()))
	fmt.Println(FormatTime(time.Now()))
	fmt.Println(TimeStamp()/(1000*60*60*24))
	fmt.Println(time.Now().Unix()/(60*60*24))
	fmt.Println(time.Now().Hour())*/
	timesta := time.Now().Unix()
	fmt.Println(timesta+GetBetweenNextDayByCurrentTime(timesta))
	fmt.Println(GetBetweenNextDayTimeByTime(timesta).Unix())
	fmt.Println(GetBetweenNextDayTimeByTime2(timesta))
	/*timeStr := time.Now().Format("2006-01-02")
	fmt.Println(timeStr)

	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)

	fmt.Println(t.Unix() + 1)
	fmt.Println(t2.AddDate(0, 0, 1).Unix())*/
}


func Zipmap(a, b []string) (map[string]string, error) {

	if len(a) != len(b) {
		return nil, errors.New("zip: arguments must be of same length")
	}

	r := make(map[string]string, len(a))

	for i, e := range a {
		r[e] = b[i]
	}

	return r, nil
}

//随机数
func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

//生成32位md5值
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//格式化时间
func FormatTime(t time.Time) string {
	f := "2006-01-02 15:04:05"
	return t.Format(f)
}

//字符串转换为时间戳
func Stime2TimeStamp(stime string) int64 {
	f := "2006-01-02 15:04:05"
	t, e := time.ParseInLocation(f, stime, time.Local)
	if e != nil {
	}
	return t.UnixNano() / 1e6
}

//获取当前时间戳，为毫秒
func TimeStamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//获取当前时间戳，为毫秒
func UTCTimeStamp() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}

//创建token
func CreateToken(code, accounts, uid, time string) string {
	return Md5(code + accounts + uid + time)
}

func StrToUint(strNumber string, value interface{}) (err error) {
	var number interface{}
	number, err = strconv.ParseUint(strNumber, 10, 64)
	switch v := number.(type) {
	case uint64:
		switch d := value.(type) {
		case *uint64:
			*d = v
		case *uint:
			*d = uint(v)
		case *uint16:
			*d = uint16(v)
		case *uint32:
			*d = uint32(v)
		case *uint8:
			*d = uint8(v)
		}
	}
	return
}

func CreateCaptcha(n int) string {
	return fmt.Sprintf("%0"+fmt.Sprintf("%d", n)+"v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow(float64(10), float64(n)))))
}

//取得当前日期 20200724
func GetDate() string  {
	t := time.Now()
	return t.Format("20060102")
}

//距离下一天零晨时间间隔秒
func GetBetweenNextDay() int64 {
	currentDate := time.Now() //当前时间
	timeStr := currentDate.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	nextDayBeging := t.AddDate(0, 0, 1).Unix() //第二天时间
	return nextDayBeging - currentDate.Unix()
}

//第二天零晨时间戳，精确到秒
func GetNextDay() int64 {
	currentDate := time.Now() //当前时间
	timeStr := currentDate.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.AddDate(0, 0, 1).Unix()
}
//第二天零晨时间
func GetNextDayTime() time.Time {
	currentDate := time.Now() //当前时间
	timeStr := currentDate.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.AddDate(0, 0, 1)
}

/**
下一天开始时间
sec
	秒
*/
func GetBetweenNextDayByCurrentTime(currentTime int64) int64 {
	tm := time.Unix(currentTime, 0)
	timeStr := tm.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	nextDayBeging := t.AddDate(0, 0, 1).Unix() //第二天时间
	return nextDayBeging - currentTime
}

/**
下一天开始时间
sec
	秒
*/
func GetBetweenNextDayTimeByTime(sec int64) time.Time {
	tm := time.Unix(sec, 0)
	timeStr := tm.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return t.AddDate(0, 0, 1)
}

/**
下一天开始时间
sec
	秒
*/
func GetBetweenNextDayTimeByTime2(sec int64) int64 {
	timeSeconds := sec /// 1000
	intDaysec := timeSeconds / (24 * 3600)
	intDaysec = (intDaysec)*24*3600
	return  intDaysec - timeSeconds
}

//取得当前日期 20200724
func GetDateByTime(sec int64) string  {
	tm := time.Unix(sec, 0)
	return tm.Format("20060102")
}

func Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
