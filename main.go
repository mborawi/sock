package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/mborawi/sock/ws"
)

var upgrader = websocket.Upgrader{}

var cm ws.ConnManager

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/ws", wsHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Println("Listening on port 7070")
	go SendMessages()
	log.Fatal(http.ListenAndServe(":7070", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	cm.AddConn(conn)
}

func SendMessages() {
	ch := time.Tick(800 * time.Millisecond)
	for range ch {
		if cm.Size() < 1 {
			continue
		}
		voteCount := rand.Intn(100)
		log.Printf("Sending response to a total connection count of %d\n", cm.Size())
		response := fmt.Sprintf("Votes: %d, Time: %s", voteCount, time.Now().Format("2006-01-02T15:04:05.000Z07:00"))
		count := cm.Broadcast(websocket.TextMessage, []byte(response))
		log.Printf("%d connections sent", count)
	}
}
