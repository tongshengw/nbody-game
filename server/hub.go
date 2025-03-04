package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type ClientMsg struct {
	clientptr *Client
	msg       []byte
}

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Subhubs
	subhubs map[*SubHub]bool

	// Inbound messages from the clients.
	clientMsgs chan ClientMsg

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clientMsgs: make(chan ClientMsg),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		subhubs:    make(map[*SubHub]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.clientMsgs:
			handleClientMsg(message)
		}
	}
}

func handleClientMsg(message ClientMsg) {
	jsonMap := make(map[string]interface{})

	json.Unmarshal(message.msg, &jsonMap)
	switch jsonMap["title"] {
	case "echo":
		s, ok := jsonMap["content"].(string)
		if jsonMap["content"] != nil && ok {
			fmt.Printf("%s\n", s)
			select {
			case message.clientptr.send <- []byte(s):
			}
		} else {
			log.Printf("%s message: %#v\n", "malformed input", jsonMap)
		}
	default:
		log.Printf("%s message: %#v\n", "unregonised message handleClientMsg", jsonMap)
	}
}
