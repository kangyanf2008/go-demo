package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := ":8080" // 监听的端口号

	// 创建一个TCP监听器，同时支持IPv4和IPv6
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s for both IPv4 and IPv6...\n", port)

	// 循环等待客户端连接
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		conn.Write(([]byte)("hello\n"))

		// 处理客户端连接
		go handleConnection(conn)
	}
}

// handleConnection 处理每个客户端连接
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		// 读取客户端发送的数据
		size, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			return
		}
		if size == -1 {
			conn.Close()
			break
		}
		if string(buffer[0:size]) == "\u0003" {
			conn.Close()
			break
		}
		// 向客户端发送响应
		//conn.Write(buffer[0:size]) // 回声（Echo）服务器，原样返回接收到的数据
		fmt.Printf("Handled connection from %s, %s \n", conn.RemoteAddr(), buffer[0:size])
	}

}
