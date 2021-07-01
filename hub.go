package main

type Hub struct {
	clients    map[string]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.nickname] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.nickname]; ok {
				delete(h.clients, client.nickname)
				close(client.queueMessage)
			}
		case message := <-h.broadcast:
			for nickname, client := range h.clients {
				if message.Nickname != nickname {
					select {
					case client.queueMessage <- message:
					default:
						close(client.queueMessage)
						delete(h.clients, client.nickname)
					}
				}
			}
		}
	}
}
