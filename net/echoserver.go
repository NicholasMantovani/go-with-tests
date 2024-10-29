package main

import (
	"fmt"
	"log"
	"net"
)

// FOR TESTING JUST RUN nc localhost 8070
func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8070")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Ensure we teardown the server when the program exits
	defer listener.Close()

	fmt.Println("Server is listening on port 8070")

	for {
		// Block until we receive an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Ensure we close the connection after we're done
	defer conn.Close()

	fmt.Println("new client connected", conn.RemoteAddr())
	conn.Write([]byte("Hello from the echo server!\n"))

	for {
		// Read data
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		log.Println("Received data", buf[:n])

		// Write the same data back
		conn.Write(buf[:n])
	}

}
