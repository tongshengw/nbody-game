package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

type gameStateSerialised struct {
	Player1 string `json:"player1"`
}

type playerSerialised struct {
	X float32
	Y float32
	Z float32
}
