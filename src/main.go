package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"pcap"
)

func main() {
	// 获取设备列表
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// 选择要捕获的设备
	device := devices[0].Name
	log.Printf("Using device: %s\n", device)

	// 打开设备
	handle, err := pcap.OpenLive(device, 65536, true, 10*1000)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// 设置过滤器，只捕获TCP流量
	filter := "tcp"
	err = handle.SetFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	// 开始捕获数据包
	log.Println("Starting capture...")
	loop := func(hdr *pcap.PacketHdr, pkt []byte) {
		// 处理数据包
		processPacket(handle, hdr, pkt)
	}
	handle.Loop(loop, nil)
}

func processPacket(handle *pcap.Handle, hdr *pcap.PacketHdr, pkt []byte) {
	// 在这里实现你自己的数据包分析逻辑
	fmt.Printf("Received packet at %s len=%d\n", hdr.Timestamp, hdr.Len)
}
