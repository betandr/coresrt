package receiver

import (
	"fmt"
	"log"
	"net"
)

func handlePacket(data []byte, addr *net.UDPAddr) {
	fmt.Printf("Received %d bytes from %s\n", len(data), addr.String())
}

func Start(port int, ipAddr string) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ipAddr),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("error listening on UDP port: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Listening on port %d...\n", addr.Port)
	buf := make([]byte, 2048)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error reading UDP: %v", err)
			continue
		}
		go handlePacket(buf[:n], remoteAddr)
	}
}
