package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)


type room struct {
	forward chan []byte   // 保持到来的消息
	join chan *client     // 希望加入房间的客户端
	leave chan *client    // 希望离开房间的客户端
	clients map[*client]bool // 保持房间中所有客户端
}

func (r *room) run() {
	for {
		select {
		case client := <- r.join:
			r.clients[client] = true
		case client := <- r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <- r.forward:
			for client := range r.clients {
				select {
				case client.send <- msg:
					// send msg
				default:
					// 发送失败
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: messageBufferSize}

func (r *room) ServeHttp(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)  // 升级 http -> websocket 协议
	if err != nil {
		log.Fatal("Serve Err", err)
		return
	}
	client := &client{
		socket: socket,
		send: make(chan []byte, messageBufferSize),
		room: r,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
