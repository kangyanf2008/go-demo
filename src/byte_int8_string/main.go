package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*var a byte = 122
	fmt.Println(string(a))
	fmt.Println(Test(int8(a)))
	var cc int = 122
	ee := int8(cc)
	fmt.Println(ee)
	fmt.Println(byte(ee))
	fmt.Println(string(ee))*/
	aa := "1:2"
	tem := strings.Split(aa, ":")
	fmt.Println(tem[0])
	ee,_:= strconv.Atoi(tem[0])
	fmt.Println(uint8(ee))
	fmt.Println(byte(ee))



}

func Test(s int8) string  {
	return strconv.Itoa(int(s))
}