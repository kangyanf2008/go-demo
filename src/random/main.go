package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i:= 0; i< 20; i++ {
		fmt.Println(rand.Intn(3))
	}
	fmt.Println("======================================")
	t := []byte{0,1}
	switch t[0] {
	case 0 ,3 , 4:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("fii")
	}

}
