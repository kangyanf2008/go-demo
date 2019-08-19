package main

import "fmt"

const INT_MAX = int64(^uint64(0) >> 1)

func main() {
	fmt.Println(INT_MAX)
}
