package main

import (
	"fmt"
	"os"
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

	params := os.Args
	if len(params) == 0 {
		fmt.Println("command error")
		os.Exit(-1)
	}
	paramString := strings.Join(params,"")
	filterString := []string{"exec", "inspect", "attach", "export", "cp", "save","help"}
	for _, v := range filterString {
		fmt.Println(strings.Index(strings.ToLower(paramString), strings.ToLower(v)))
		if strings.Index(strings.ToLower(paramString), strings.ToLower(v)) > 0 {
			fmt.Println("command error")
			os.Exit(-1)
		}
	}


}
