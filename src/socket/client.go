package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var clientID string
	flag.StringVar(&clientID, "clientId", "", "--clientId")
	flag.Parse()
	conn, err := net.Dial("tcp", "127.0.0.1:44444")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	conn.Write([]byte("clientId:" + clientID))
	for {
		buffer := make([]byte, 1024)
		num, readErr := conn.Read(buffer)
		if readErr != nil {
			fmt.Println(readErr)
			break
		}
		if num > 0 {
			tem := string(buffer[0:num])
			fmt.Println(tem)
			if strings.Index(tem, "disconnect") != -1 {
				break
			}
		}
	}

}
