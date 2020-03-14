package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrades = websocket.Upgrader{}

func reverse(w http.ResponseWriter, r *http.Request) {
	ws, _ := upgrades.Upgrade(w, r, nil)
	defer ws.Close()

	for {
		_, message,_ := ws.ReadMessage()
		log.Printf("Message received: %s", message)
	}
}

func main() {
	fmt.Println("Starttin websocket server al :5555")
	http.HandleFunc("/reverse", reverse)
	log.Fatal(http.ListenAndServe(":5555", nil))
}
