package main

import "fmt"

func main() {
	var result []string
	result = append(result,"11")
	result = append(result,"22")
	result = append(result,"33")
	fmt.Println(len(result))
	fmt.Println(cap(result))
	fmt.Println(result)

	var result2 []string
	fmt.Println(len(result2))
	fmt.Println(result2)
	fmt.Println(result2==nil)
	fmt.Println("截取长度")
	fmt.Println(result[0:3])


}
