package main

import ()

type clientConnection struct {
	clientptr *Client

	privateChan chan []byte
}

// Needs to be capable of:
// - recieving input from two clients
// - sending private output to each client
// - register a client
// - when hub gets given a code, spawn a subhub with that client
type SubHub struct {
	client1 clientConnection

	client2 clientConnection

	broadcast chan []byte

	unregister chan *Client

	gameptr *Game
}

func newSubHub(client1 *Client, client2 *Client) *SubHub {
	return &SubHub{
		broadcast:  make(chan []byte),
		unregister: make(chan *Client),
		client1:    clientConnection{clientptr: client1, privateChan: make(chan []byte)},
		client2:    clientConnection{clientptr: client2, privateChan: make(chan []byte)},
	}
}

func (sh *SubHub) run() {
	for {
		select {}
	}
}

