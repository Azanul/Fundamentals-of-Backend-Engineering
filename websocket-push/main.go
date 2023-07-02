package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type message struct {
	origin      net.Addr
	messageType int
	body        []byte
}

type hub struct {
	connections      map[*websocket.Conn]bool
	broadcastChannel chan message
}

func (h *hub) run() {
	for {
		msg := <-h.broadcastChannel
		for conn := range h.connections {
			if msg.origin == conn.RemoteAddr() {
				continue
			}
			h.write(conn, msg)
		}
	}
}

func main() {
	wsHub := &hub{
		make(map[*websocket.Conn]bool, 0),
		make(chan message),
	}
	go wsHub.run()

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			log.Println(err)
			return
		}

		wsHub.connections[ws] = true
		wsHub.read(ws)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h *hub) read(conn *websocket.Conn) {
	h.broadcastChannel <- message{messageType: websocket.TextMessage, body: []byte("new client joined hub")}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		h.broadcastChannel <- message{origin: conn.RemoteAddr(), messageType: messageType, body: p}
	}
}

func (h *hub) write(conn *websocket.Conn, msg message) {
	if err := conn.WriteMessage(msg.messageType, msg.body); err != nil {
		log.Println(err)
	} else {
		delete(h.connections, conn)
	}
}
