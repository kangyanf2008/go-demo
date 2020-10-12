package once_pakage2

import "once_factory"

var aa = func() {
	once_factory.Factory()
}

func init() {
	once_factory.Factory()

}



func DemoFunc2()  {
	once_factory.Factory()
}
