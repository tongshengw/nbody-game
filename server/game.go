package main

import (
	"log"
	"time"

	"github.com/go-gl/mathgl/mgl32"
)

type Game struct {
	state  GameState
	input1 chan GameInput
	subhub *SubHub
}

type Player struct {
	p mgl32.Vec3
	v mgl32.Vec3
	r mgl32.Quat
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
		input1: make(chan GameInput),

		state: GameState{
			player1: Player{
				p: mgl32.Vec3{0, 0, 0},
				v: mgl32.Vec3{0, 0, 0},
				r: mgl32.QuatIdent(),
			},
		},

		subhub: sh,
	}
}

func (g *Game) run() {
	const TPS float32 = 64
	ticker := time.NewTicker(time.Second / time.Duration(TPS))
	defer ticker.Stop()

	// game loop
	for {
		<-ticker.C

		// process input
		g.processInput()

		// update velocities
		// gravity physics not implemented yet

		// calculate positions
		g.state.player1.p = g.state.player1.p.Add(g.state.player1.v.Mul(1 / TPS))

		// output
		g.subhub.client1.send <- g.state.toserial()
		// log.Printf("%.2f, %.2f\n", g.state.player1.p.X(), g.state.player1.p.Z())

	}
}

func (g *Game) processInput() {
	firstInputProcessed := false
	for {
		select {
		case pin, ok := <-g.input1:
			if ok && !firstInputProcessed {
				firstInputProcessed = true
				var velocityChange = mgl32.Vec3{0, 0, 0}
				if pin.D_pressed {
					velocityChange = velocityChange.Add(mgl32.Vec3{1, 0, 0})
				}
				if pin.S_pressed {
					velocityChange = velocityChange.Add(mgl32.Vec3{0, 0, 1})
				}
				if pin.A_pressed {
					velocityChange = velocityChange.Add(mgl32.Vec3{-1, 0, 0})
				}
				if pin.W_pressed {
					velocityChange = velocityChange.Add(mgl32.Vec3{0, 0, -1})
				}
				g.state.player1.v = g.state.player1.v.Add(velocityChange)
			}
			// discard everything not processed
		default:
			return
		}
	}
}
