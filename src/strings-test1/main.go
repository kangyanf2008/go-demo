package main

import (
	"fmt"
	"strings"
)

func main() {
	t1 := strings.Split("+86|185000000000","|")
	fmt.Println(t1)
	t2 := strings.Split("5555,666",",")
	fmt.Println(t2)
	t3 := strings.Split("5555666",",")
	fmt.Println(t3)
	aa := "/songshu/track/overseas_idn_shuaji/20190823/songshu_track-overseas_idn_shuaji.log.1.1.1.18.20190822010000"
	fmt.Println(aa[1:len(aa)])
	fmt.Println(strings.ReplaceAll(aa[1:len(aa)],"/","_"))
}
