package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		linesChan := getLinesChannel(conn)

		for line := range linesChan {
			fmt.Println(line)
		}
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
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}
			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}
			currentLineContents += parts[len(parts)-1]
		}
	}()
	return lines
}// package main

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	//	"log"
// 	"net"
// 	"os"
// 	"strings"
// )

// const inputFilePath = "messages.txt"

// func main() {
// 	listener, err := net.Listen("tcp", ":42069")
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server is listening on port 42069")
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		fmt.Printf("conn type is: %T\n\n",conn)
// 		fmt.Println("Connection accepted...")
// 		lines := getLinesChannel(conn)
// 		for line := range lines {
// 			fmt.Println(line)
// 		}
// 		fmt.Println("Connection has been closed.")
// 		//		go handleConnection(conn)
// 	}

// }

// func getLinesChannel(f io.ReadCloser) <-chan string {
// 	lines := make(chan string)
// 	go func() {
// 		defer f.Close()
// 		defer close(lines)
// 		currentLineContents := ""
// 		for {
// 			b := make([]byte, 8, 8)
// 			n, err := f.Read(b)
// 			if err != nil {
// 				if currentLineContents != "" {
// 					lines <- currentLineContents
// 				}
// 				if errors.Is(err, io.EOF) {
// 					break
// 				}
// 				fmt.Printf("error: %s\n", err.Error())
// 				return
// 			}
// 			str := string(b[:n])
// 			parts := strings.Split(str, "\n")
// 			for i := 0; i < len(parts)-1; i++ {
// 				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
// 				currentLineContents = ""
// 			}
// 			currentLineContents += parts[len(parts)-1]
// 		}
// 	}()
// 	return lines
// }


// package main

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// const inputFilePath = "messages.txt"

// func main() {
// 	f, err := os.Open(inputFilePath)
// 	if err != nil {
// 		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
// 	}
// 	defer f.Close()

// 	fmt.Printf("Reading data from %s\n", inputFilePath)
// 	fmt.Println("=====================================")
// 	for line := range getLinesChannel(f) {
// 		fmt.Printf("read: %s\n", line)
// 	}
// }

// func getLinesChannel(f io.ReadCloser) <-chan string {
// 	out := make(chan string, 1)

// 	go func() {
// 		defer f.Close()
// 		defer close(out)
// 		currentLineContents := ""
// 		for {
// 			buffer := make([]byte, 8, 8)
// 			n, err := f.Read(buffer)
// 			if err != nil {
// 				if currentLineContents != "" {
// 					out <- currentLineContents
// 					currentLineContents = ""
// 				}
// 				if errors.Is(err, io.EOF) {
// 					break
// 				}
// 				fmt.Printf("error: %s\n", err.Error())
// 				break
// 			}
// 			str := string(buffer[:n])
// 			parts := strings.Split(str, "\n")
// 			for i := 0; i < len(parts)-1; i++ {
// 				currentLineContents += parts[i]
// 				out <- currentLineContents
// 				currentLineContents = ""
// 			}
// 			currentLineContents += parts[len(parts)-1]
// 		}
// 	}()

// 	return out
// }
