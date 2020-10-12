package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int ,100 )
	for i := 0; i< 20 ; i ++ {
		queue<-i
	}
//	close(queue)
//	fmt.Println(11)
	queue<-11
//	fmt.Println(22)
	close(queue)

	go func() {
		for {
			select  {
			case s, e := <-queue:
				if e {
					fmt.Println("g1",s, e)
					time.Sleep(time.Second)
				} else {
					break
				}
			}
		}
		fmt.Println("g2 exit")
	}()

	go func() {
		for {
			select  {
			case s, e := <-queue:
				if e {
					fmt.Println("g2",s, e)
					time.Sleep(time.Second)
				} else {
					break
				}
			}
		}
		fmt.Println("g2 exit")
	}()

	select {
	}

}
