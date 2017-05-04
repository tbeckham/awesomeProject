//
// Hello World Zeromq Client
//
// Author: Tony Beckham   github.com/tbeckham
// adapted from the Go example from http://zguide.zeromq.org/go:hwserver
// Requires: http://github.com/pebbe/zmq4
//
package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"time"
)

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()
	socket.Bind("tcp://*:5555")

	// Wait for messages
	for {
		msg, _ := socket.Recv(0)
		println("Received ", string(msg))

		// do some fake "work"
		time.Sleep(time.Second)

		// send reply back to client
		reply := fmt.Sprintf("World")
		socket.SendBytes([]byte(reply), 0)
	}
}