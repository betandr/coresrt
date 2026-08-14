package receiver

import (
	"log"
	"net"
	"sync"
	"time"
)

type Options struct{}

type Receiver struct {
	conn        *net.UDPConn
	connections map[string]*connection // key: remote addr string
	mu          sync.Mutex
	startTime   time.Time
}

type connection struct {
	socketID   uint32
	peerSocket uint32
	addr       *net.UDPAddr
	cookie     uint32
	startTime  time.Time

	// Sequence tracking for ACKs
	mu              sync.Mutex
	lastAckedSeq    uint32 // last sequence number we ACKed
	highestSeq      uint32 // highest contiguous sequence number received
	seqInitialized  bool   // whether we've seen the first data packet
	ackNumber       uint32 // sequential ACK number (starts at 1)
	packetsReceived uint64 // total packets received
	bytesReceived   uint64 // total bytes received
	firstPacketTime time.Time
	lastPacketTime  time.Time
	connected       bool // true after handshake complete
	stopACK         chan struct{}
}

func Start(port int, ipAddr string, opts Options) {
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ipAddr),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("error listening on UDP port: %v", err)
	}
	defer conn.Close()

	// r := &Receiver{
	// 	conn:        conn,
	// 	connections: make(map[string]*connection),
	// 	startTime:   time.Now(),
	// }

	log.Printf("SRT listener started on %s:%d\n", ipAddr, port)
	log.Printf("waiting for SRT connections...")

	buf := make([]byte, 2048)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("error reading UDP: %v", err)
			continue
		}

		// copy packet to use with shared buffer
		pktData := make([]byte, n)
		copy(pktData, buf[:n])

		log.Println("received packet from", remoteAddr.String(), "size:", n)
		log.Printf("first 32 bytes: % x", pktData[:min(32, len(pktData))])

		// pkt, err := packets.ParsePacket(pktData)
		// if err != nil {
		// 	log.Printf("[%s] error parsing packet: %v", remoteAddr.String(), err)
		// 	continue
		// }

		// switch p := pkt.(type) {
		// case *packets.Data:
		// 	r.handleDataPacket(p, remoteAddr)
		// case *packets.Control:
		// 	go r.handleControlPacket(p, remoteAddr)
		// }
	}
}
