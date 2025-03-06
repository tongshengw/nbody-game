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
	unallocatedClients map[*Client]bool

	// Subhubs, clients map to subhubs because each subhub has one client at a time for now
	subhubs map[*Client]*SubHub

	// Inbound messages from the clients.
	clientMsgs chan ClientMsg

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clientMsgs:         make(chan ClientMsg),
		register:           make(chan *Client),
		unregister:         make(chan *Client),
		unallocatedClients: make(map[*Client]bool),
		subhubs:            make(map[*Client]*SubHub),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.unallocatedClients[client] = true
		case client := <-h.unregister:
			if _, ok := h.unallocatedClients[client]; ok {
				delete(h.unallocatedClients, client)
				close(client.send)
			}

		case message := <-h.clientMsgs:
			handleClientMsg(h, message)
		}
	}
}

func (message *ClientMsg) marshalToMap() map[string]interface{} {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(message.msg, &jsonMap)
	if err != nil {
		log.Printf("marshalToMap error: %s\n", message.msg)
	}
	_, ok := jsonMap["title"]
	if !ok {
		log.Printf("title malformed\n")
	}
	return jsonMap
}

func handleClientMsg(h *Hub, message ClientMsg) {
	jsonMap := message.marshalToMap()
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
	case "request_subhub":
		_, ok := h.subhubs[message.clientptr]
		if ok {
			log.Printf("%s\n", "subhubs map already contains clientptr")
		} else {
			newSubHubPtr := newSubHub(message.clientptr)
			h.subhubs[message.clientptr] = newSubHubPtr
			message.clientptr.subhub = newSubHubPtr
			fmt.Printf("%s\n", "subhub request recieved")
			go newSubHubPtr.run()
		}

	default:
		log.Printf("%s message: %#v\n", "unregonised message handleClientMsg", jsonMap)
	}
}
