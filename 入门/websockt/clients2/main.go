package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}
func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://localhost:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}
