package main

import (
	// "fmt"
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
	ch := time.Tick(500 * time.Millisecond)
	res := payload{}
	for range ch {
		if cm.Size() < 1 {
			continue
		}
		res.Sydney += rand.Float64() * 5
		res.Melbourne += rand.Float64() * 4.3
		res.Canberra += rand.Float64() * 6
		res.Newcastle += rand.Float64() * 3

		// log.Printf("Sending response to a total connection count of %d\n", cm.Size())
		// response := fmt.Sprintf("%.2f", sum)
		count := cm.BroadcastJson(res)
		log.Printf("%d connections sent", count)
		if res.Sydney > 100 {
			res.Sydney = 0
		}
		if res.Melbourne > 100 {
			res.Melbourne = 0
		}
		if res.Canberra > 100 {
			res.Canberra = 0
		}
		if res.Newcastle > 100 {
			res.Newcastle = 0
		}
	}
}

type payload struct {
	Sydney    float64 `json:"sydney"`
	Melbourne float64 `json:"melbourne"`
	Newcastle float64 `json:"newcastle"`
	Canberra  float64 `json:"canberra"`
}
