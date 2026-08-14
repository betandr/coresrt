package receiver

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strings"
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

		dumpPacket("SRT packet from "+remoteAddr.String(), pktData)

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

func dumpPacket(label string, pktData []byte) {
	if len(pktData) == 0 {
		log.Printf("%s: empty packet", label)
		return
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s (%d bytes):\n", label, len(pktData)))
	b.WriteString(packetGroups(pktData))
	b.WriteString("\n")

	for i := 0; i < len(pktData); i += 16 {
		end := i + 16
		if end > len(pktData) {
			end = len(pktData)
		}

		chunk := pktData[i:end]
		b.WriteString(fmt.Sprintf("%04x: ", i))

		for _, v := range chunk {
			b.WriteString(fmt.Sprintf("%02x ", v))
		}
		for len(chunk) < 16 {
			b.WriteString("   ")
			chunk = nil
		}
		b.WriteString("  ")

		for _, v := range pktData[i:end] {
			if v >= 32 && v < 127 {
				b.WriteByte(byte(v))
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteString("\n")
	}

	log.Print(b.String())
}

func packetGroups(pktData []byte) string {
	if len(pktData) < 16 {
		return "packet too short for SRT header\n"
	}

	var b strings.Builder
	b.WriteString("header groups:\n")

	first := pktData[0]
	if first&0x80 != 0 {
		ctlFlag := (first >> 7) & 0x01
		ctlType := uint16(first&0x7F)<<8 | uint16(pktData[1])
		subType := uint16(pktData[2])<<8 | uint16(pktData[3])
		typeSpecific := binary.BigEndian.Uint32(pktData[4:8])
		timestamp := binary.BigEndian.Uint32(pktData[8:12])
		destSocket := binary.BigEndian.Uint32(pktData[12:16])

		b.WriteString(fmt.Sprintf("  control flag: %d\n", ctlFlag))
		b.WriteString(fmt.Sprintf("  control type: %015b (%04x)\n", ctlType, ctlType))
		b.WriteString(fmt.Sprintf("  subtype: %016b (%04x)\n", subType, subType))
		b.WriteString(fmt.Sprintf("  type-specific info: %032b (%08x)\n", typeSpecific, typeSpecific))
		b.WriteString(fmt.Sprintf("  timestamp: %032b (%08x)\n", timestamp, timestamp))
		b.WriteString(fmt.Sprintf("  destination socket id: %032b (%08x)\n", destSocket, destSocket))
	} else {
		seqNum := binary.BigEndian.Uint32([]byte{pktData[0] & 0x7F, pktData[1], pktData[2], pktData[3]})
		msgFlags := pktData[4]
		msgNum := uint32(pktData[4]&0x3F)<<18 | uint32(pktData[5])<<10 | uint32(pktData[6])<<2 | uint32(pktData[7]>>6)
		timestamp := binary.BigEndian.Uint32(pktData[8:12])
		destSocket := binary.BigEndian.Uint32(pktData[12:16])

		b.WriteString(fmt.Sprintf("  packet seq num: %031b (%08x)\n", seqNum, seqNum))
		b.WriteString(fmt.Sprintf("  flags: %08b\n", msgFlags))
		b.WriteString(fmt.Sprintf("  message num: %026b (%08x)\n", msgNum, msgNum))
		b.WriteString(fmt.Sprintf("  timestamp: %032b (%08x)\n", timestamp, timestamp))
		b.WriteString(fmt.Sprintf("  destination socket id: %032b (%08x)\n", destSocket, destSocket))
	}

	return b.String()
}
