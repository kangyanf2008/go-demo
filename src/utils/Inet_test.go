package utils

import (
	"testing"
)

func TestINet_AToN(t *testing.T) {
	ip := "192.168.10.1"
	ipv4Num, err := INet_AToN(ip)
	if err != nil {
		t.Fatalf("INet_AToN(%s)执行失败, err=%v", ip, err)
		t.Fail()
	}
	t.Logf("INet_AToN(%s)执行成功, ipv4Num=%d", ip, ipv4Num)
	ipv4String, err := INet_NToA(ipv4Num)
	if err != nil {
		t.Errorf("INet_NToA(%d)执行失败, err=%v", ipv4Num, err)
		t.Fail()
	}
	if ipv4String != ip {
		t.Errorf("INet_NToA(%d)结果不正确, 期望值=%v,实际值=%v", ipv4Num, ip, ipv4String)
		t.Fail()
	}
	t.Logf("INet_NToA(%d)正确, 值=%v", ipv4Num, ipv4String)
}
