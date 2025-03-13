package main

import (
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	state  GameState
	inputs chan GameInput
	subhub *SubHub
}

type Player struct {
	position mgl32.Vec3
}

type GameInput struct {
	Player    int
	W_pressed bool `json:"w_pressed"`
	A_pressed bool `json:"a_pressed"`
	S_pressed bool `json:"s_pressed"`
	D_pressed bool `json:"d_pressed"`
}

type GameState struct {
	player1 Player
}

func newGame(sh *SubHub) *Game {
	return &Game{
		inputs: make(chan GameInput),
		state:  GameState{player1: Player{position: mgl32.Vec3{0, 0, 0}}},
		subhub: sh,
	}
}

func (g *Game) run() {
	const TPS = 64
	ticker := time.NewTicker(time.Second / time.Duration(TPS))
	defer ticker.Stop()

	// game loop
	var pos float32 = 0.01
	for {
		<-ticker.C
		// log.Printf("%s", g.state.toserial())
		g.subhub.client1.send <- g.state.toserial()
		g.state.player1.position[0] += pos
		g.state.player1.position[1] += pos * 2
		g.state.player1.position[2] += pos * 3
		// process input
		// update velocities
		// calculate positions
		// output
	}
}
