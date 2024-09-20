package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server at localhost:8080")
	reader := bufio.NewReader(os.Stdin)
	serverScanner := bufio.NewScanner(conn)

	for {
		fmt.Print("Enter message (type 'exit' to quit): ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		message = message[:len(message)-1]

		if message == "exit" {
			fmt.Println("Closing connection.")
			break
		}

		_, err = fmt.Fprintf(conn, message+"\n")
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		if serverScanner.Scan() {
			response := serverScanner.Text()
			fmt.Println("Server response:", response)
		} else {
			if err := serverScanner.Err(); err != nil {
				fmt.Println("Error reading server response:", err)
			} else {
				fmt.Println("Server closed the connection.")
			}
			break
		}
	}
}
