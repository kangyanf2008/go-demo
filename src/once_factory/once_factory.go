package once_factory

import (
	"fmt"
	"sync"
)

var once sync.Once

func init() {
	fmt.Println("init1")
/*	once.Do(func() {
		fmt.Println("11111111")
	})*/
}

func Factory()  {
	fmt.Println("factory")
}