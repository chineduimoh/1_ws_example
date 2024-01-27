package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
		"syscall"

	"github.com/gorilla/websocket"
)

var connCount int = 0

func ws(w http.ResponseWriter, r *http.Request) {
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

	// Read messages from socket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Client Disconnected: ", connCount)
			connCount--
			conn.Close()
			break
		}
		log.Printf("msg: %s", string(msg))
	}
}

func main() {
		// Set Ulimit
	var rlimit syscall.Rlimit

	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		fmt.Println(err)
	}

	rlimit.Cur = rlimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Ulimit set to maximum")
	
	http.HandleFunc("/", ws)
	fmt.Println("Server started & listening on port: ", os.Getenv("PORT"))
	if err := http.ListenAndServe(":" + os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
