package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	conn, err := net.Listen("tcp", "0.0.0.0:44444")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		os.Exit(1)
	}
	// 获取操作系统的文件描述符
	connectPool := new(sync.Map)

	defer func() {
		conn.Close()

		connectPool.Range(func(key, value any) bool {
			clientConnect := value.(*net.TCPConn)
			clientConnect.Write([]byte(key.(string) + ",disconnect"))
			clientConnect.Close()
			connectPool.Delete(key)
			return true
		})
	}()

	go func() {
		for {
			fmt.Println("begin check client connect")
			connectPool.Range(func(key, value any) bool {
				clientConnect := value.(*net.TCPConn)
				fmt.Println("连接池", key, clientConnect.RemoteAddr())
				_, writeErr := clientConnect.Write([]byte(key.(string) + ",heartbeat!!"))
				if writeErr != nil {
					connectPool.Delete(key)
				}
				return true
			})
			fmt.Println("end check client connect")
			time.Sleep(time.Second * 5)
		}
	}()
	for {
		newConnect, accErr := conn.Accept()
		if accErr != nil {
			fmt.Println(accErr)
			break
		}

		go func() {
			clientIdBuffer := make([]byte, 1024)
			readNum, readErr := newConnect.Read(clientIdBuffer)
			if readErr != nil {
				newConnect.Close()
				return
			}
			var clientId string
			if readNum > 0 {
				clientId = string(clientIdBuffer[0:readNum])
			}

			fdFile, _ := newConnect.(*net.TCPConn).File()
			fd := fdFile.Fd()
			key := fmt.Sprintf("%s,fd:%v", clientId, fd)
			connectPool.Store(key, newConnect)
			fmt.Println(key, "ipAddr:", newConnect.RemoteAddr())

			defer func() {
				connectPool.Delete(key)
				newConnect.Close()
			}()

			i := 0
			for i < 100 {
				i++
				_, writeErr := newConnect.Write([]byte(clientId + ", hello:" + strconv.Itoa(i)))
				if writeErr != nil {
					break
				}
				time.Sleep(time.Second)
			}
			newConnect.Write([]byte("disconnect!"))
		}()

	}

}
