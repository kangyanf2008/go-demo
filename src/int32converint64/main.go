package main

import (
	"fmt"
	"strconv"
)

func main() {
	 var aa int64 = 1234567891

	 var bb int32
	 bb = int32(aa)
	 fmt.Println(bb)
	 dd := "9876543210"
	 userid,_ := strconv.ParseInt(dd,10,32)
	 fmt.Println(int32(userid))
	 nn, _ := strconv.ParseInt(dd,10,64)
	 fmt.Println(nn)

	birthday := "0"
	if len(birthday) > 0 {
		birthdayInt64, _ := strconv.ParseInt(birthday, 10, 64)
		fmt.Println(birthdayInt64)
	}

	var a int = 1000000000
	fmt.Println(int32(a))

	var cc int32 = 1000000000
  have := false
	fmt.Println(int(cc))
have = true
fmt.Println(have)

var test int64 =  10000
fmt.Println(test)
fmt.Println(int32(test))

}
