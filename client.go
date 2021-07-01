package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Nickname string `json:"nickname,omitempty"`
	Content  string `json:"content,omitempty"`
}

type Client struct {
	nickname     string
	hub          *Hub
	conn         *websocket.Conn
	queueMessage chan Message
}

func (c *Client) readWS() {}

func (c *Client) writeWS() {}

func handleWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query()["nickname"]

	if len(nickname) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("cannot get the websocket connection: ", err)
		return
	}

	client := &Client{
		nickname:     nickname[0],
		hub:          hub,
		conn:         conn,
		queueMessage: make(chan Message, 2),
	}

	client.hub.register <- client
	go client.writeWS()
	go client.readWS()
}
