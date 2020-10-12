package once_pakege3

import "once_factory"
var AA = func() {
	once_factory.Factory()
}


func init() {
	once_factory.Factory()
	AA()
}

func init(){
	once_factory.Factory()
}

func init(){
	once_factory.Factory()
}

func DemoFunc2()  {
	go once_factory.Factory()
}

