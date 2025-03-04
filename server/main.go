package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

	fmt.Printf("%s\nnbody-server\n", nbody_ascii)

	var addr = flag.String("addr", ":8080", "http service address")
	flag.Parse()
	fmt.Printf("Serving on: %s\n", *addr)

	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHomePage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "serverHome.html")
}