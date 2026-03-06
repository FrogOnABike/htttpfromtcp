package main

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
	