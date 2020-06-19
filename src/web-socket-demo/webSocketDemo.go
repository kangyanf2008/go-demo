package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"sync"
)

var (
	pool sync.Map
)

func main() {
	http.Handle("/ws", websocket.Handler(wsHandler))
	if err := http.ListenAndServe(":9999",nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func wsHandler(conn *websocket.Conn)  {
	for {
		if !conn.IsServerConn() && !conn.IsClientConn() {
			return
		}
		fmt.Println(conn.Request().Header)
		data := ""
		if err := websocket.Message.Receive(conn, &data); err != nil {
			if err.Error() == "EOF" {
				fmt.Println("client close read")
				return
			}
			fmt.Println("read error=",err)
			continue
		}
		if err := websocket.Message.Send(conn, data); err != nil {
			if err.Error() == "EOF" {
				fmt.Println("client close write")
				return
			}
			fmt.Println("write error=",err)
			continue
		}
		fmt.Println(data)
	}
}