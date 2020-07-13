package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wait := &sync.WaitGroup{}
	//ctx, cancel := context.WithTimeout(context.Background(),time.Second*10)
	ctx, cancel := context.WithCancel(context.Background())
	runNum := new(int)
	exitNum := new(int)

	for i:=0; i< 100000; i++  {
		wait.Add(1)
		go worker(ctx, wait, runNum, exitNum)
	}

	//休眠退出
	go func() {
		time.Sleep(time.Second*10)
		cancel()
	}()

	//等待协程执行完毕后退出
	wait.Wait()
	fmt.Println("exit sub goroutine ,runNum=", *runNum, ",exitNum=", *exitNum)
	fmt.Println("main goroutine exit")
}

func worker(ctx context.Context, w *sync.WaitGroup, run, exit *int)  {
	*run++
	defer func() {
		*exit++
		w.Done()
	}()
	LABLE:
	for  {
		select {
		case <-ctx.Done():
			break LABLE
		default:
			//fmt.Println("run go id=",getGID(),"\exit")
		}
	}
	fmt.Println("go id=",getGID()," exit \n")
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}