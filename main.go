package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
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
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")
	//  L4 code below: still read in 8 byte chunks but print each complete line, starting a new line for each line read.
	//  Each new line should start with "read:"
	// If the last line does not end with a newline character, print it as well.

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
