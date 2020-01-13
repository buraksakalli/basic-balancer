package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int
	listenAddr  = "localhost:8080"

	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal("failed to listen: %s", err)
	}
	
	defer listener.Close()
	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
		}

		backend := chooseBackend()
		fmt.Println("counter=%d backend=%s\n", counter, backend)
		go func(){
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("Warning: Proxying failed %v", err)
			}
		}()
	}
}

func proxy(backend string, connection net.Conn) error {
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed %s: %v", backend, err)
	}

	//backend to connection
	go io.Copy(bc, connection)

	// connection to backend
	go io.Copy(connection, bc)

	return nil
}

func chooseBackend() string {
	choosedServer := server[counter%len(server)]
	counter++
	return choosedServer
}