package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func promptUsername(conn *websocket.Conn) {
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte(
			"Please enter a username: \n"))
		if err != nil {
			log.Println(err)
		}

		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		promptUsername(conn)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))

}
