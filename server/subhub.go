package main

import ()

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

	unregister chan *Client

	gameptr *Game
}

func newSubHub(client1 *Client) *SubHub {
	return &SubHub{
		unregister: make(chan *Client),
		client1:    client1,
		// client2:    client2,
	}
}

func (sh *SubHub) run() {
	for {
		select {}
	}
}

