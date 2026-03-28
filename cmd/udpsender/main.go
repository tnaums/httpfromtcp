package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Resolve the server address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Error resolving server address:", err)
	}

	// Establish UDP connection to server
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('c')
		if err != nil {
			fmt.Println("Error reading string:", err)
			return
		}

		// Send message to server
		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("Error sending message: %v", err)
			continue
		}
		fmt.Printf("Message sent: %s", line)
	}

}
