// dcgsscmd project main.go
package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// we accept all origins
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	log.Println("drone course ground station (dcgss) started")
	// handle websocket
	http.HandleFunc("/ws", wshandler)
	// start listening on port 8505
	log.Fatal(http.ListenAndServe(":8505", nil))
}

// websocket handler: /ws?port=14550&protocol=udp
func wshandler(w http.ResponseWriter, r *http.Request) {
	protocol := r.URL.Query().Get("protocol")
	port := r.URL.Query().Get("port")

	// accepts only udp protocol to connect to the drone
	if protocol != "udp" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start listening on the given port
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start listening
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer ServerConn.Close()

	// buffer to read the incoming data from the udp connection
	buf := make([]byte, 1024)

	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		// read incomming MAVLink message from udp
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			return
		}

		// write them to the opened websocket
		if err := conn.WriteMessage(websocket.BinaryMessage, buf[0:n]); err != nil {
			log.Println(err)
			return
		}
	}
}
