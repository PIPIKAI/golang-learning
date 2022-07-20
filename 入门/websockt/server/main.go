package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// websocket 相当于升级版的http
var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conns []*websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	conns = append(conns, conn)
	for {
		m, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, conn := range conns {
			conn.WriteMessage(websocket.TextMessage, []byte("你说的是："+string(p)+" 吗？"))
		}
		fmt.Println(m, string(p))
	}
	defer conn.Close()
	log.Println("服务关闭")
}
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Fprintf(os.Stderr, "ERR :main Err: %v", err)
	}
}
