package model

type InvokeConn struct {
	PInfo PrivideInfo
	ObjectPool interface{}
	ObjectPoolConfig interface{}
}


func New(pInfo PrivideInfo, objectPool interface{}, objectPoolConfig interface{}) *InvokeConn {
	ic := new(InvokeConn)
	ic.ObjectPool = objectPool
	ic.PInfo = pInfo
	ic.ObjectPoolConfig = objectPoolConfig
	return  ic
}