package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Needs to be capable of:
// - recieving input from two clients
// - sending private output to each client
// - register a client
// - when hub gets given a code, spawn a subhub with that client
type SubHub struct {
	client1 *Client

	// client2 *Client

	unregister chan *Client

	clientMsgs chan ClientMsg

	gameptr *Game
}

type TitleOnly struct {
	Title string `json:"title"`
}

type PlayerInputMsg struct {
	Title string    `json:"title"`
	Data  GameInput `json:"data"`
}

func newSubHub(client1 *Client) *SubHub {
	return &SubHub{
		client1:    client1,
		clientMsgs: make(chan ClientMsg),
		unregister: make(chan *Client),
		gameptr:    nil,
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

	paused := false
	for !paused {
		select {
		case <-sh.unregister:
			// TODO: add logic to close subhubs and games that have been open too long
			paused = true

		case message := <-sh.clientMsgs:
			shHandleClientMsg(sh, &message)
		}
	}
}

func shHandleClientMsg(sh *SubHub, message *ClientMsg) {
	fmt.Printf("subhub message: %s\n", message.msg)
	var titleStruct TitleOnly
	json.Unmarshal(message.msg, &titleStruct)
	switch titleStruct.Title {
	case "game_ready":
		sh.gameptr = newGame(sh)
		go sh.gameptr.run()
	case "player_input_data":
		if sh.gameptr != nil {
			var playerIn GameInput
			json.Unmarshal(message.msg, &playerIn)
			playerIn.Player = 1
			sh.gameptr.inputs <- playerIn
		}
	}
}
