package main

import "fmt"

var str *[]string
var str2 []string

func main() {
	setVal(str)
	setVal2(&str2)
	println(str)
	println(&str2)

	println(str)
	println(&str2)
	println(str2[0], str2[1])

	d := &Demo{}
	setValue(d)
	fmt.Println(d)
	fmt.Println(*d)
	fmt.Println(&d)

}

//需要在这里赋值str，但是又不能直接引用 str
func setVal(val *[]string) {
	val = &[]string{"a", "b"}
}

func setVal2(val *[]string) {
	*val = []string{"a", "b"}
}


type Demo struct {
	name string
	age int
}

func setValue(d *Demo)  {
	d.name = "aa"
	d.age = 111
}