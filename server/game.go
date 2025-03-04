package main

import ()

type Game struct {
	state  *gameState
	inputs chan GameInput
}

type player struct {
}

type GameInput struct {
	player int
}

type gameState struct {
	player1 player
}

func newGame() *Game {
	return nil
}
