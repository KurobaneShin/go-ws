package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("New connection from %s\n", conn.RemoteAddr().String())

	// Sample data to send
	data := []byte{0x01, 0x02, 0x03, 0x04, 0xFF, 0xFE, 0xFD, 0x00}

	for _, b := range data {
		_, err := conn.Write([]byte{b})
		if err != nil {
			fmt.Printf("Error sending data: %v\n", err)
			return
		}

		fmt.Printf("Sent byte: %02x\n", b)

		// Add a small delay between sends for demonstration purposes
		time.Sleep(100 * time.Millisecond)

		if b == 0x00 {
			fmt.Println("Sent 0x00, closing connection")
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}
