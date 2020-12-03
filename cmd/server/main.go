package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Cannot establish TCP connection!")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 100)

	for {
		conn.Write([]byte("Hi! I am your web-server, what is your name?"))

		lenData, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message: ", string(buf))

		conn.Write(append([]byte("Nice to meet you "), buf[:lenData]...))

	}
}
