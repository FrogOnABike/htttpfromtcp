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

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %s\n", inputFilePath, err)
	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")
	//  L4 code below: still read in 8 byte chunks but print each complete line, starting a new line for each line read.
	//  Each new line should start with "read:"
	// If the last line does not end with a newline character, print it as well.

	var line string
	for {
		b := make([]byte, 8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if line != "" {
					fmt.Printf("read: %s\n", line)
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
				fmt.Printf("read: %s\n", line)
				line = part
			}
		}
	}
}

// L3 code below: read 8 bytes at a time and print the string representation of the bytes read
// for {
// 	b := make([]byte, 8, 8)
// 	n, err := f.Read(b)
// 	if err != nil {
// 		if errors.Is(err, io.EOF) {
// 			break
// 		}
// 		fmt.Printf("error: %s\n", err.Error())
// 		break
// 	}
// 	str := string(b[:n])
// 	fmt.Printf("read: %s\n", str)
// }
