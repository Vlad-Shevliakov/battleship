package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Cannot establish TCP connection!")
		return
	}

	go copyTo(os.Stdout, conn)

	sendMessage(conn)

}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func sendMessage(conn net.Conn) {

	f := os.Stdin

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		conn.Write(scanner.Bytes())
	}
}

func printMessage(conn net.Conn) {

	message := make([]byte, 100)

	_, err := conn.Read(message)
	if err != nil {
		fmt.Println("failed to read message")
	}

	fmt.Println("> ", string(message))

}
