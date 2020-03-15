package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrades = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func reader(conn *websocket.Conn) {
	for {
		_, message, _ := conn.ReadMessage()
		log.Printf("Message received: %s", message)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	upgrades.CheckOrigin = func(r *http.Request) bool {return true}

	ws, _ := upgrades.Upgrade(w, r, nil)
	defer ws.Close()

	reader(ws)
}

func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "server is Alive")
}

func setupRoutes() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", alive)
}

func main() {
	fmt.Println("Starting websocket server al :5555")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":5555", nil))
}
