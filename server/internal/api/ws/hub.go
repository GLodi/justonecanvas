package ws

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client

	log *logrus.Logger
}

func NewHub(l *logrus.Logger) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		log:        l,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.log.Infoln("hub + connected:", len(h.clients))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				h.log.Infoln("hub - connected:", len(h.clients))
			}
		case message := <-h.broadcast:
			// HACK: this is needed for artillery testing
			// message[0] = message[0] - 48
			// message[1] = message[1] - 48
			// message[2] = message[2] - 48
			h.log.Infoln("hub received:", message)

			// go doEvery(20*time.Millisecond, helloworld, h)
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func doEvery(d time.Duration, f func(time.Time, *Hub), hub *Hub) {
	for x := range time.Tick(d) {
		f(x, hub)
	}
}

func helloworld(t time.Time, hub *Hub) {
	message := make([]byte, 3)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := r1.Intn(16)
	b := r1.Intn(40)
	c := r1.Intn(40)
	// hub.log.Infoln("a", byte(a))
	// hub.log.Infoln("b", byte(b))
	// hub.log.Infoln("b", byte(c))
	message[0] = byte(a)
	message[1] = byte(b)
	message[2] = byte(c)

	for client := range hub.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(hub.clients, client)
		}
	}
}
