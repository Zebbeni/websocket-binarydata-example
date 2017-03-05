package main

import (
	"fmt"
	"flag"
	"net/http"
	"golang.org/x/net/websocket"
	"time"
	"encoding/binary"
)



var port *int = flag.Int("port", 1234, "Port to listen")
var root *string = flag.String("root", ".", "root to serve")


var trigger_channels = map[chan bool]bool{}

var add_chan = make(chan chan bool)
var del_chan = make(chan chan bool)

var bindata = []byte{0,0,0,0, 0,0,0,0}  // This is the byte array to be sent through websocket

func add_trigger_channel(chan2 chan chan bool) {
	for newchan := range chan2 {
		trigger_channels[newchan] = true
	}
}

func del_trigger_channel(chan2 chan chan bool) {
	for newchan := range chan2 {
		trigger_channels[newchan] = false
	}
}


func gendata() {  // This function generates/updates data and puts in bindata variable.
	counter := uint64(1024)
	for {
		// fill bindata
		binary.LittleEndian.PutUint64(bindata, counter)
		//fmt.Println(bindata)
		
		// Now send notifications to all dataFeeder go routines
		for c, active := range trigger_channels {
			if active == true {
				c <- true
			}
		}
		counter += uint64(1)
		time.Sleep(time.Second*2)
	}
	
}

func main() {
	flag.Parse()

	go add_trigger_channel(add_chan)
	go del_trigger_channel(del_chan)
	go gendata()
	
	http.Handle("/", http.FileServer(http.Dir("build/web/")))
	http.Handle("/tabledata", websocket.Handler(dataFeeder))
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
	return
}


func dataFeeder(ws *websocket.Conn) {
	fmt.Println("Received a ws connection")
	trigger_channel := make(chan bool, 100)
	add_chan <- trigger_channel
	fmt.Println("Added trigger_channel")

	for _ = range trigger_channel {
		fmt.Println("data-generator updated bindata.")
		err := websocket.Message.Send(ws, bindata)
		fmt.Println("Send returned err = ", err)
		if err != nil {  // Client closed conn ?
			del_chan <- trigger_channel
			break
		}
	}
	return
}
