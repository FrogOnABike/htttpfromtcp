package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const inputFilePath = "messages.txt"

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)
	go func() {
		defer close(lines)
		defer f.Close()
		var line string
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if line != "" {
						lines <- line
					}
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				break
			}
			parts := strings.Split(string(b[:n]), "\n")
			for i, part := range parts {
				if i == 0 {
					line += part
				} else {
					lines <- line
					line = part
				}
			}
		}
	}()
	return lines
}

func main() {
	// // Open the file for reading. If the file does not exist, print an error message and exit.
	// f, err := os.Open(inputFilePath)
	// if err != nil {
	// 	log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	// }

	// fmt.Printf("Reading data from %s\n", inputFilePath)
	// fmt.Println("=====================================")

	n, err := net.Listen("tcp4", ":42069")
	if err != nil {
		log.Fatalf("could not listen: %s\n", err)
	}
	fmt.Printf("Listening on %s\n", n.Addr().String())
	defer n.Close()

	for {
		conn, err := n.Accept()
		if err != nil {
			log.Printf("could not accept connection: %s\n", err)
			continue
		}
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr().String())
		go func() {
			defer conn.Close()
			lines := getLinesChannel(conn)
			for line := range lines {
				fmt.Println(line)
			}
		}()
		fmt.Println("Connection closed")
	}
}
