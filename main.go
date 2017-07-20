// dcgsscmd project main.go
// Quickstart: ws = new WebSocket("ws://localhost:8505/ws?port=14550&protocol=udp")
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/kardianos/osext"
)

// UDP servers mapped by port	 numbers
var serverUDP map[string]*net.UDPConn

// MAVLink channels
var channelMAVLink map[string]chan []byte

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// we accept all origins
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	log.Println("drone course ground station (dcgss) started")
	log.Println("Version: " + VERSION)

	if err := checkUpdate(); err != nil {
		log.Println(err)
	} else {
		// restart
		executable, err := osext.Executable()
		if err != nil {
			log.Fatalln(err)
		}
		if err := syscall.Exec(executable, os.Args, os.Environ()); err != nil {
			log.Fatalln(err)
		}
	}

	// variable initialisation
	serverUDP = make(map[string]*net.UDPConn)
	channelMAVLink = make(map[string]chan []byte)

	// handle websocket
	http.HandleFunc("/ws", wshandler)
	//handle version
	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, VERSION)
	})
	// handle files
	http.Handle("/", http.FileServer(assetFS()))

	// open browser
	go browse("localhost:8505")

	// start listening on port 8505
	log.Fatal(http.ListenAndServe(":8505", nil))
}

func browse(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "http://localhost:8505/").Start()
	default:
		log.Println("Go to http://localhost:8505/ to access server")
	}
	return err
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

	if _, ok := serverUDP[port]; ok == false {
		// ServerConn does not exist, we create it and add it to the map

		// get random address to listen to on the given port
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

		// add connection to the map
		serverUDP[port] = ServerConn

		// create channel and add to the map
		MAVLinkChannel := make(chan []byte, 8)
		channelMAVLink[port] = MAVLinkChannel

		// start routine
		go udpHandler(MAVLinkChannel, ServerConn)
	}

	// retrieve MAVLinkChannel from map
	MAVLinkChannel := channelMAVLink[port]

	// upgrade connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// read incoming MAVLink message from channel
	for buf := range MAVLinkChannel {
		// write them to the opened websocket
		if err := conn.WriteMessage(websocket.BinaryMessage, buf); err != nil {
			log.Println(err)
			return
		}
	}
}

func udpHandler(ch chan<- []byte, ServerConn *net.UDPConn) {
	// buffer to read the incoming data from the udp connection
	buf := make([]byte, 1024)
	warned := false
	for {
		// read incomming MAVLink message from udp
		n, _, err := ServerConn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			return
		}

		// send to channel
		select {
		case ch <- buf[0:n]:
		default:
			if warned == false {
				log.Println("Channel is full, message has been discarded.")
				warned = true
			}
		}
	}
}
