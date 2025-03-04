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
	
	var addr = flag.String("addr", ":8080", "http service address")
	
	func serveHome(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		http.ServeFile(w, r, "home.html")
	}
	
	func main() {
		flag.Parse()
		hub := newHub()
		go hub.run()
		http.HandleFunc("/", serveHome)
		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			serveWs(hub, w, r)
		})
		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
