package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	var filename = "D:\\mobile.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}

	if err1 != nil {
		fmt.Println(err1)
		return
	}
	for i:=0; i< 9000; i++ {
		num := ""
		if i < 10 {
			num += ("000"+strconv.Itoa(i))
		}
		if i >= 10 && i < 100 {
			num += ("00"+strconv.Itoa(i))
		}
		if i >= 100 && i < 1000 {
			num += ("0"+strconv.Itoa(i))
		}
		if i >= 1000 && i < 10000 {
			num += strconv.Itoa(i)
		}
		n, err1 := io.WriteString(f,     "    \"+861700000"+num+"\"=\"1111\" \n") //写入文件(字符串)
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(n)
	}

	defer f.Close()
}


/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}