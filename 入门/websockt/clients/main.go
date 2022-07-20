package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

var reader = bufio.NewReader(os.Stdin)

func send(conn *websocket.Conn, username string) {

	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, []byte(username+","+string(l)))
	}
}

type Client struct {
	userName string
	urlStr   string
}

func NewClient() *Client {
	fmt.Println("输入客户端的用户名")
	usname, _, _ := reader.ReadLine()
	return &Client{
		userName: string(usname),
		urlStr:   "ws://localhost:8888",
	}
}
func (c *Client) Run() {
	fmt.Printf("```开始通讯```\n Url：%s \n 用户名：%s", c.urlStr, c.userName)
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial(c.urlStr, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go send(conn, c.userName)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(m, string(p))
	}
}
func main() {
	c := NewClient()
	c.Run()
}
