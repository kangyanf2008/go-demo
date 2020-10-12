package once_factory

import (
	"fmt"
)


func init() {
	fmt.Println("init2")
/*	once.Do(func() {
		fmt.Println("11111111")
	})*/
}

func Factory2()  {
	fmt.Println("factory")
}