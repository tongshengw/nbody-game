package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func main() {
	nbody_ascii := `
      ___.              .___      
  ____\_ |__   ____   __| _/__.__.
 /    \| __ \ /  _ \ / __ <   |  |
|   |  \ \_\ (  <_> ) /_/ |\___  |
|___|  /___  /\____/\____ |/ ____|
     \/    \/            \/\/
	`

	fmt.Printf("%s\nStarting nbody-server...\n", nbody_ascii)
}
