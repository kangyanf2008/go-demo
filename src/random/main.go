package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
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
fmt.Println("=============================")
	connects := make(map[int64]*Connect_info)
	connects[11]=&Connect_info{Ip:"111"}
	aa ,ok := connects[11]
	fmt.Println(aa)
	fmt.Println(ok)

	ee ,ll := connects[22]
	fmt.Println(ll)
	fmt.Println(ee)

	testStr := "\\0"+"aaaaa"
	fmt.Println(testStr)


	var intUserid int64 = 0

	intUserid=11111111
	defer Print(intUserid)
	time.Sleep(time.Second)
}

func Print(u int64 )  {
	fmt.Println(u)
}

type Connect_info struct {
	Conn           net.Conn //连接
	Keepalive_time int64    //纳秒
	Ip             string   //客户端Ip
}