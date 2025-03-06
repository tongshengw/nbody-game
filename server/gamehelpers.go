package main

import (
	"encoding/json"
	"log"
)

type gameStateSerialised struct {
	Title   string `json:"title"`
	Player1 string `json:"player1"`
}

type playerSerialised struct {
	Xpos float32 `json:"posx"`
	Ypos float32 `json:"posy"`
	Zpos float32 `json:"posz"`
}

func (p *Player) toserial() []byte {
	pSerial := playerSerialised{Xpos: p.position.X(), Ypos: p.position.Y(), Zpos: p.position.Z()}
	val, err := json.Marshal(pSerial)
	if err != nil {
		log.Printf("player toserial() json marshal error")
	}
	return val
}

func (gs *GameState) toserial() []byte {
	gsSerial := gameStateSerialised{Title: "gamestate", Player1: string(gs.player1.toserial())}
	val, err := json.Marshal(gsSerial)
	if err != nil {
		log.Printf("gamestate toserial() json marshal error")
	}
	return val
}
