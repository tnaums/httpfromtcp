package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/tnaums/httpfromtcp/internal/request"
)

const port = 42069

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())					
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
		log.Fatalf("error: %s\n", err.Error())
			continue
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		//linesChan := getLinesChannel(conn)
		r, _ := request.RequestFromReader(conn)
		
		fmt.Printf("%s\n", r)
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)
		currentLineContents := ""
		for {
			b := make([]byte, 8, 8)
			n, err := f.Read(b)
			currentLineContents += string(b[:n])
			parts := strings.SplitN(currentLineContents, "\n", 2)
			if len(parts) > 1 {
				lines <- parts[0]
				currentLineContents = parts[1]
			}
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}

		}
		lines <- currentLineContents
	}()
	return lines
}
