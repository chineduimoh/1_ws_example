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
	fmt.Println("Ulimit set to maximum", rlimit.Max)

	server := &http.Server{
		// Other configurations...
		Addr: ":9000",
	}

	// Set a custom ErrorLog
	server.ErrorLog = log.New(os.Stdout, "custom error: ", log.LstdFlags)

	http.HandleFunc("/", ws)

	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	// if err := http.ListenAndServe(":9000", nil); err != nil {
	// 	log.Fatal(err)
	// }
}
