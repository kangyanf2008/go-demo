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

}
