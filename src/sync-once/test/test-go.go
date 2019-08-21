package test

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var once sync.Once

func Init(){
	fmt.Println(GoID())
}

func Demo1(n *sync.WaitGroup)  {
	defer n.Done()
	fmt.Printf("beging demo1,gid=[%s] \n",GoID())
	once.Do(Init)
	fmt.Printf("end demo1,gid=[%s] \n",GoID())
}

func Demo2(n *sync.WaitGroup)  {
	defer n.Done()
	fmt.Printf("beging demo2,gid=[%s] \n",GoID())
	once.Do(Init)
	fmt.Printf("end demo2,gid=[%s] \n",GoID())
}

func Demo3(n *sync.WaitGroup)  {
	defer n.Done()
	fmt.Printf("beging demo3,gid=[%s] \n",GoID())
	once.Do(Init)
	fmt.Printf("end demo3,gid=[%s] \n",GoID())
}

//获取协和ID
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
