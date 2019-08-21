package main

import (
	"sync"
	"sync-once/test"
)

func main() {
	var go_sync sync.WaitGroup
	for i:=0; i< 10; i++ {
		go test.Demo1(&go_sync)
		go_sync.Add(1)
		go  test.Demo2(&go_sync)
		go_sync.Add(1)
		go  test.Demo3(&go_sync)
		go_sync.Add(1)
	}
	go_sync.Wait()
}
