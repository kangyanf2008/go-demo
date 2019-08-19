package loadbalance

import (
	"balance/model"
	"balance/service_map"
	"fmt"
)
const INT_MAX = int(^uint(0) >> 1)

//var current = service_map.NewSafeMap()
var current = service_map.NewSafe2Map()

type RoundRobinLoadBalanceStrategy struct {}

func NewRoundRobinLoadBalanceStrategy(obj map[string]int) *RoundRobinLoadBalanceStrategy {
	strategy := new(RoundRobinLoadBalanceStrategy)
	for k, v := range obj {
		current.WriteMap(k,v)
	}
	return strategy
}

func (strategy RoundRobinLoadBalanceStrategy) DoSelect(invokeConns []model.InvokeConn, invocation model.Invocation, callId string) {
	key := invocation.InterfaceName + invocation.MethodName
	cur := strategy.getMapValue(key)
	fmt.Println(invokeConns[(cur % len(invokeConns))].PInfo, cur, callId)
}

func (strategy RoundRobinLoadBalanceStrategy) getMapValue( key string) int {
	current.Lock()
	defer current.Unlock()
	cur := current.ReadMap(key)
	if INT_MAX -1  <= cur{
		cur = 0
	}
	current.WriteMap(key, cur + 1)
	return cur
}

func (strategy RoundRobinLoadBalanceStrategy) GetCallNum(invocation model.Invocation,) int {
	key := invocation.InterfaceName + invocation.MethodName
	return current.ReadMap(key)
}