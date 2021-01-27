package main

import (
	"fmt"
	"net"
	"time"
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
		conn.Write([]byte("\n"))
		conn.Write([]byte("Hi! I am your web-server, what is your name?\n"))

		lenData, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Message: ", string(buf))

		resp := fmt.Sprintf("Nice to meet you, %v!\n", string(buf[:lenData]))

		conn.Write(append([]byte(resp)))
		time.Sleep(3 * time.Second)
	}
}
