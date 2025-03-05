package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

type Game struct {
	state  *GameState
	inputs chan GameInput
}

type player struct {
	position mgl32.Vec3
}

type GameInput struct {
	player int
}

type GameState struct {
	player1 player
}

func newGame() *Game {
	return nil
}

func (g Game) run() {
	const TPS = 64
	ticker := time.NewTicker(time.Second / time.Duration(TPS))
	defer ticker.Stop()

	// game loop
	for {
		<-ticker.C

	}
}
