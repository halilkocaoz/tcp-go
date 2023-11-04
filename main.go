package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", ":2319")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		netConnection, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleNetConnection(netConnection)
	}
}

func handleNetConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr())
	for {
		received, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Printf("%s: %s\n", c.RemoteAddr(), err)
			break
		}

		if received == "CLOSE\n" {
			break
		}

		response := []byte("Message received: " + received)
		_, err = c.Write(response)
		if err != nil {
			fmt.Println(err)
		}
	}
	c.Close()
	fmt.Printf("Closing %s\n", c.RemoteAddr())
}
