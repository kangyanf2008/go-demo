package main

import "fmt"

func main() {
	keys := make([]int64, 0, 5)
	fmt.Println(len(keys))
	fmt.Println(cap(keys))
	keys = append(keys, 1)
	fmt.Println(len(keys))
	fmt.Println(cap(keys))
}
