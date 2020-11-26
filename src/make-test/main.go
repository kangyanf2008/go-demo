package main

import (
	"fmt"
)

func main() {
/*	keys := make([]int64, 0, 5)
	fmt.Println(len(keys))
	fmt.Println(cap(keys))
	keys = append(keys, 1)
	fmt.Println(len(keys))
	fmt.Println(cap(keys))
*/
	//fmt.Println(int(math.Ceil(float64(1) / float64(1))))
	var aa string
	keys := map[string]string{"aa":"eeeeeeeeee","ee":"22"}
	fmt.Println(keys)
	if ee, ok := keys["aa"]; ok {
		aa = ee
		fmt.Println(aa)
	}
	fmt.Println("------------"+aa)
}
