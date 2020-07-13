package main

import (
	"balance/loadbalance"
	"balance/model"
	"fmt"
	"time"
)

func main() {
	obj := make(map[string]int)
	obj["login.aaa"]=0
	obj["login.out"]=0
	strategy := loadbalance.NewRoundRobinLoadBalanceStrategy(obj)
	conns := make([]model.InvokeConn, 6)

	conns[0].PInfo = model.PrivideInfo{Ip:"192.168.1.1", Port:9090}
	conns[1].PInfo = model.PrivideInfo{Ip:"192.168.1.2", Port:9090}
	conns[2].PInfo = model.PrivideInfo{Ip:"192.168.1.3", Port:9090}
	conns[3].PInfo = model.PrivideInfo{Ip:"192.168.1.4", Port:9090}
	conns[4].PInfo = model.PrivideInfo{Ip:"192.168.1.5", Port:9090}
	conns[5].PInfo = model.PrivideInfo{Ip:"192.168.1.6", Port:9090}

	inte1 := model.Invocation{MethodName:"aaa",InterfaceName:"login."}
	//inte2 := model.Invocation{MethodName:"out",InterfaceName:"login."}
	for i:=0; i< 100000; i++{
		strategy.DoSelect(conns, inte1, fmt.Sprint(i))
		//go strategy.DoSelect(conns, inte1, fmt.Sprint(i))
		//go strategy.DoSelect(conns, inte1, fmt.Sprint(i))
		//go strategy.DoSelect(conns, inte2)
	}
	time.Sleep(time.Second*2)
	fmt.Println(strategy.GetCallNum(inte1))
}