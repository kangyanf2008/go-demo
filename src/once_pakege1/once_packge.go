package once_pakege1

import "once_factory"

var aa = func() {
	once_factory.Factory()
}

func init() {
	once_factory.Factory()

}

func init(){
	once_factory.Factory()
}



func DemoFunc2()  {
	go once_factory.Factory()
}

