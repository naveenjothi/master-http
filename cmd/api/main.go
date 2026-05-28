package main

import (
	"fmt"
	"io"
	"net"
)

const port = "8080"

func buildHTTPResponse(body string) string {
	return fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n"+
			"Content-Type: text/plain\r\n"+
			"Content-Length: %d\r\n"+
			"\r\n"+
			"%s",
		len(body),
		body,
	)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Connection from: %s\n", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	// for {
	n, err := conn.Read(buffer)

	if err != nil {
		if err == io.EOF {
			fmt.Printf("Connection closed by client: %s\n", conn.RemoteAddr())
		} else {
			fmt.Println("Error reading:", err)
		}
		return
	}

	request := string(buffer[:n])

	fmt.Println("RAW REQUEST:")
	fmt.Println(request)

	body := "Hello World"

	response := buildHTTPResponse(body)

	_, err = conn.Write([]byte(response))

	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	// }
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
