package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
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

	ws, err := upgrades.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close()

	reader(ws)
}

func alive(w http.ResponseWriter, r *http.Request) {
	//t:= template.NewTemplateHAndler("../../templates/chat.html")
	//t.ServeHTTP(w,r)
	fmt.Fprintf(w, "Home page")
}

func setupRoutes() {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", alive)
}

var doOnce sync.Once

func startWebSocketServer(port string) {
	doOnce.Do(func(){
		log.Fatal(http.ListenAndServe(":" + port, nil))
	})
}


func main() {
	fmt.Println("Starting websocket server al :5555")
	setupRoutes()
	startWebSocketServer("5555")
}
