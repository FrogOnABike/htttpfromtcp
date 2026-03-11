package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "localhost:42069")
	if err != nil {
		log.Fatalf("could not resolve address: %s\n", err)
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatalf("could not dial: %s\n", err)
	}
	defer conn.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		text, err := input.ReadString('\n')
		if err != nil {
			log.Printf("could not read input: %s\n", err)
			continue
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Printf("could not send message: %s\n", err)
		}
	}
}
