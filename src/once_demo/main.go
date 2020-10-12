package main

import (
	"once_factory"
	once_pakege12 "once_pakege1"
	once_pakege1 "one_package3"
	"time"
)
var aa = func() {
	once_factory.Factory()
}

func init() {
	once_factory.Factory()

}

func init(){
	once_factory.Factory()
}

func init(){
	once_factory.Factory()
}


func main() {
	go func() {
		once_factory.Factory()
		go once_pakege1.DemoFunc2()
		go once_pakege1.DemoFunc2()
		go once_pakege12.DemoFunc2()
	}()
	time.Sleep(time.Second*2)

}
