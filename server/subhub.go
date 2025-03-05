package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ClientConnection struct {
	clientptr *Client

	privateChan chan []byte
}

// Needs to be capable of:
// - recieving input from two clients
// - sending private output to each client
// - register a client
// - when hub gets given a code, spawn a subhub with that client
type SubHub struct {
	client1 *Client

	// client2 *Client

	clientMsgs chan ClientMsg

	gameptr *Game
}

func newSubHub(client1 *Client) *SubHub {
	return &SubHub{
		client1:    client1,
		clientMsgs: make(chan ClientMsg),
		gameptr:    newGame(),
		// client2:    client2,
	}
}

func (sh *SubHub) run() {
	type msgConfirmSubhub struct {
		Title string `json:"title"`
	}
	msg, err := json.Marshal(msgConfirmSubhub{Title: "subhub_allocated"})
	if err != nil {
		log.Printf("confirm subhub json marshal error")
	} else {
		sh.client1.send <- msg
	}

	for {
		select {
		case message := <-sh.clientMsgs:
			shHandleClientMsg(sh, &message)
		}
	}
}

func shHandleClientMsg(sh *SubHub, message *ClientMsg) {
	fmt.Printf("subhub message: %s\n", message.msg)
	jsonMap := message.marshalToMap()
	switch jsonMap["title"] {
	case "game_ready":
		sh.gameptr = newGame()
		go sh.gameptr.run()
	}
}
