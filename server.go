package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var connCount int = 0

func ws(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i got here")
	// Upgrade connection
	connCount++
	if connCount%100 == 0 {
		fmt.Println("Client Connected: ", connCount)
	}

	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// Read messages from socket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Client Disconnected: ", connCount)

			return
		}
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
	http.HandleFunc("/", ws)
	fmt.Println("Server started")
	if err := http.ListenAndServe(":" + os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
