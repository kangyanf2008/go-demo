package loadbalance

import "balance/model"

type ILoadBalanceStrategy interface {

	DoSelect( invokeConns []model.InvokeConn, invocation model.Invocation, callId string)
}