package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listnener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting TCP server", err)
		return
	}

	defer listnener.Close()
	fmt.Println("Server is listening on port 8000")

	for {
		conn, err := listnener.Accept()
		if err != nil {
			fmt.Println("Error accessing connection", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connected to", conn.RemoteAddr())

	//read data from conn
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("recieved: ", message)

		_, err := conn.Write([]byte("Echo: " + message + "\n"))
		if err != nil {
			fmt.Println("error writing to client", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Connection error: ", err)
	}

	fmt.Println("Connection closed: ", conn.RemoteAddr())
}
