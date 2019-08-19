package model

type ServcieInfo struct {
	version      string  //服务版本号
	group        string  //服务集群组名
	directInvoke bool    //是否只是直接调用.true表示直接调用,false表示会额外注册到zk上.
}
