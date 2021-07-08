package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Stream struct {
	Date  string  `json:"date"`
	Time  string  `json:"time"`
	Price float64 `json:"price"`
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("err: ", err)
		return
	}
	defer c.Close()

	// st := "yeah"
	for i := 0; i < 5; i++ {
		// err := c.WriteMessage(websocket.TextMessage, []byte(st))
		st := Stream{
			time.Now().Format("2006-01-02"),
			time.Now().Format("15:04:05"),
			float64(i)*0.01 + 1,
		}
		err := c.WriteJSON(st)
		if err != nil {
			log.Print("err: ", err)
			break
		} else {
			log.Print(st)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./js")))
	http.HandleFunc("/ws", echo)

	log.Println("server starting...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
