package main

import (
	"fmt"
	"net"
)

type TCPConn struct {
	net.Conn
}

func NewTCPConn(addr string) *TCPConn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to", addr)
	return &TCPConn{
		Conn: conn,
	}
}

func (conn *TCPConn) ReadLoop(dataChan chan<- byte, errChan chan<- error) {
	for {
		data := make([]byte, 1)
		n, err := conn.Read(data)
		if err != nil {
			errChan <- err
			return
		}
		for i := 0; i < n; i++ {
			dataChan <- data[i]
		}
	}
}

func main() {
	conn := NewTCPConn("localhost:3000")
	defer conn.Close()

	dataChan := make(chan byte)
	errChan := make(chan error)

	go conn.ReadLoop(dataChan, errChan)

	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Received byte: %02x\n", data)
			if data == 0x00 {
				fmt.Println("Done")
				return
			}
		case err := <-errChan:
			fmt.Println("Error", err)
		}
	}
}
