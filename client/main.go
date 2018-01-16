package main

import (
	"bufio"
	"flag"
  "github.com/bkkwok/go-cli-chat/event"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	} else {
		log.Print("You have successfully connected to cli-chat at ", *addr)
	}

	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer conn.Close()
		defer close(done)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read: ", err)
				return
			}
			log.Printf("recv: %s", msg)
		}
	}()

	go func() {
		defer conn.Close()
		defer close(done)
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			err := conn.WriteMessage(websocket.TextMessage, []byte(text))
			if err != nil {
				log.Println("write: ", err)
			}

		}
	}()

	for {
		<-interrupt
		log.Println("interrupt")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
		err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}

		<-done
		conn.Close()
	}
}
