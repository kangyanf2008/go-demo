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
}
