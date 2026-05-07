package main

import (
	"fmt"
	"net"
)

const port = "8080"

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received:", message)

	response := "Hello World\n"

	_, err = conn.Write([]byte(response))

	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server listening on port %s", port)

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}
