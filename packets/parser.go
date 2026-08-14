package packets

import (
	"encoding/binary"
	"fmt"
	"net"
)

// MinPacketSize is the minimum size of an SRT packet (16 bytes header).
const MinPacketSize = 16

// HandshakeCIFSize is the fixed size of the handshake CIF (without extensions).
// 4 (version) + 2 (enc) + 2 (ext) + 4 (ISN) + 4 (MTU) + 4 (flow) + 4 (type) + 4 (socket) + 4 (cookie) + 16 (IP) = 48
const HandshakeCIFSize = 48

// IsControlPacket checks if the first bit of the packet is set,
// indicating a control packet.
func IsControlPacket(data []byte) bool {
	return data[0]&0x80 != 0
}

// ParsePacket parses raw bytes into either a Data or Control packet.
// It returns the parsed packet and an error if parsing fails.
func ParsePacket(data []byte) (interface{}, error) {
	if len(data) < MinPacketSize {
		return nil, fmt.Errorf("packet too short: %d bytes (minimum %d)", len(data), MinPacketSize)
	}

	if IsControlPacket(data) {
		return ParseControlPacket(data)
	}
	return ParseDataPacket(data)
}

// ParseDataPacket parses raw bytes into a Data packet struct.
func ParseDataPacket(data []byte) (*Data, error) {
	if len(data) < MinPacketSize {
		return nil, fmt.Errorf("data packet too short: %d bytes", len(data))
	}

	pkt := &Data{}

	// First 4 bytes: 0|Packet Sequence Number (31 bits)
	firstWord := binary.BigEndian.Uint32(data[0:4])
	pkt.PacketSequenceNumber = firstWord & 0x7FFFFFFF

	// Second 4 bytes: PP(2)|O(1)|KK(2)|R(1)|Message Number(26)
	secondWord := binary.BigEndian.Uint32(data[4:8])
	pkt.PacketPositionFlag = byte((secondWord >> 30) & 0x03)
	pkt.OrderFlag = byte((secondWord >> 29) & 0x01)
	pkt.KeyBasedEncryptionFlag = byte((secondWord >> 27) & 0x03)
	pkt.RetransmittedPacketFlag = byte((secondWord >> 26) & 0x01)
	pkt.MessageNumber = secondWord & 0x03FFFFFF

	// Bytes 8-11: Timestamp
	pkt.Timestamp = binary.BigEndian.Uint32(data[8:12])

	// Bytes 12-15: Destination Socket ID
	pkt.DestinationSocketID = binary.BigEndian.Uint32(data[12:16])

	// Remaining bytes: Data payload
	if len(data) > MinPacketSize {
		pkt.Data = make([]byte, len(data)-MinPacketSize)
		copy(pkt.Data, data[MinPacketSize:])
	}

	return pkt, nil
}

// ParseControlPacket parses raw bytes into a Control packet struct.
func ParseControlPacket(data []byte) (*Control, error) {
	if len(data) < MinPacketSize {
		return nil, fmt.Errorf("control packet too short: %d bytes", len(data))
	}

	pkt := &Control{}

	// First 4 bytes: 1|Control Type(15)|Subtype(16)
	firstWord := binary.BigEndian.Uint32(data[0:4])
	pkt.ControlType = ControlPacketType((firstWord >> 16) & 0x7FFF)
	pkt.Subtype = ControlPacketType(firstWord & 0xFFFF)

	// Bytes 4-7: Type-specific Information
	pkt.TypeSpecificInfo = binary.BigEndian.Uint32(data[4:8])

	// Bytes 8-11: Timestamp
	pkt.Timestamp = binary.BigEndian.Uint32(data[8:12])

	// Bytes 12-15: Destination Socket ID
	pkt.DestinationSocketID = binary.BigEndian.Uint32(data[12:16])

	// Remaining bytes: Control Information Field (CIF)
	if len(data) > MinPacketSize {
		pkt.ControlInformationField = make([]byte, len(data)-MinPacketSize)
		copy(pkt.ControlInformationField, data[MinPacketSize:])
	}

	return pkt, nil
}

// ControlTypeName returns a human-readable name for a control packet type.
func ControlTypeName(ct ControlPacketType) string {
	switch ct {
	case HANDSHAKE:
		return "HANDSHAKE"
	case KEEPALIVE:
		return "KEEPALIVE"
	case ACK:
		return "ACK"
	case NAK:
		return "NAK"
	case CongestionWarning:
		return "CONGESTION_WARNING"
	case SHUTDOWN:
		return "SHUTDOWN"
	case ACKACK:
		return "ACKACK"
	case DROPREQ:
		return "DROPREQ"
	case PEERERROR:
		return "PEERERROR"
	case UserDefinedType:
		return "USER_DEFINED"
	default:
		return fmt.Sprintf("UNKNOWN(0x%04X)", uint16(ct))
	}
}

// HandshakeTypeName returns a human-readable name for a handshake type.
func HandshakeTypeName(ht HandshakeType) string {
	switch ht {
	case Done:
		return "DONE"
	case Agreement:
		return "AGREEMENT"
	case Conclusion:
		return "CONCLUSION"
	case WaveHand:
		return "WAVEHAND"
	case Induction:
		return "INDUCTION"
	default:
		return fmt.Sprintf("UNKNOWN(0x%08X)", uint32(ht))
	}
}

// ParseHandshakeCIF parses the Control Information Field of a HANDSHAKE packet.
func ParseHandshakeCIF(cif []byte) (*HandshakeControl, error) {
	if len(cif) < HandshakeCIFSize {
		return nil, fmt.Errorf("handshake CIF too short: %d bytes (minimum %d)", len(cif), HandshakeCIFSize)
	}

	hs := &HandshakeControl{}
	hs.Version = binary.BigEndian.Uint32(cif[0:4])
	hs.EncryptionField = CypherFamilyAndKeySize(binary.BigEndian.Uint16(cif[4:6]))
	hs.ExtensionField = HandshakeExtensionFlag(binary.BigEndian.Uint16(cif[6:8]))
	hs.InitialPacketSequenceNumber = binary.BigEndian.Uint32(cif[8:12])
	hs.MaximumTransmissionUnitSize = binary.BigEndian.Uint32(cif[12:16])
	hs.MaximumFlowWindowSize = binary.BigEndian.Uint32(cif[16:20])
	hs.HandshakeType = HandshakeType(binary.BigEndian.Uint32(cif[20:24]))
	hs.SRTSocketID = binary.BigEndian.Uint32(cif[24:28])
	hs.SYNCookie = binary.BigEndian.Uint32(cif[28:32])
	hs.PeerIPAddress = PeerIPAddress{
		IP1: binary.BigEndian.Uint32(cif[32:36]),
		IP2: binary.BigEndian.Uint32(cif[36:40]),
		IP3: binary.BigEndian.Uint32(cif[40:44]),
		IP4: binary.BigEndian.Uint32(cif[44:48]),
	}

	// Parse extensions if present
	if len(cif) > HandshakeCIFSize {
		remaining := cif[HandshakeCIFSize:]
		if len(remaining) >= 4 {
			hs.ExtensionType = ExtensionType(binary.BigEndian.Uint16(remaining[0:2]))
			hs.ExtensionLength = binary.BigEndian.Uint16(remaining[2:4])
			extBytes := int(hs.ExtensionLength) * 4
			if len(remaining) >= 4+extBytes {
				hs.ExtensionContents = make([]byte, extBytes)
				copy(hs.ExtensionContents, remaining[4:4+extBytes])
			}
		}
	}

	return hs, nil
}

// SerializeHandshakeCIF serializes a HandshakeControl struct into bytes for the CIF.
func SerializeHandshakeCIF(hs *HandshakeControl) []byte {
	buf := make([]byte, HandshakeCIFSize)

	binary.BigEndian.PutUint32(buf[0:4], hs.Version)
	binary.BigEndian.PutUint16(buf[4:6], uint16(hs.EncryptionField))
	binary.BigEndian.PutUint16(buf[6:8], uint16(hs.ExtensionField))
	binary.BigEndian.PutUint32(buf[8:12], hs.InitialPacketSequenceNumber)
	binary.BigEndian.PutUint32(buf[12:16], hs.MaximumTransmissionUnitSize)
	binary.BigEndian.PutUint32(buf[16:20], hs.MaximumFlowWindowSize)
	binary.BigEndian.PutUint32(buf[20:24], uint32(hs.HandshakeType))
	binary.BigEndian.PutUint32(buf[24:28], hs.SRTSocketID)
	binary.BigEndian.PutUint32(buf[28:32], hs.SYNCookie)
	binary.BigEndian.PutUint32(buf[32:36], hs.PeerIPAddress.IP1)
	binary.BigEndian.PutUint32(buf[36:40], hs.PeerIPAddress.IP2)
	binary.BigEndian.PutUint32(buf[40:44], hs.PeerIPAddress.IP3)
	binary.BigEndian.PutUint32(buf[44:48], hs.PeerIPAddress.IP4)

	// Append extension if present
	if len(hs.ExtensionContents) > 0 {
		ext := make([]byte, 4+len(hs.ExtensionContents))
		binary.BigEndian.PutUint16(ext[0:2], uint16(hs.ExtensionType))
		binary.BigEndian.PutUint16(ext[2:4], hs.ExtensionLength)
		copy(ext[4:], hs.ExtensionContents)
		buf = append(buf, ext...)
	}

	return buf
}

// SerializeControlPacket serializes a full control packet (header + CIF).
func SerializeControlPacket(ctrl *Control) []byte {
	headerSize := MinPacketSize
	buf := make([]byte, headerSize+len(ctrl.ControlInformationField))

	// First word: 1|ControlType(15)|Subtype(16)
	firstWord := uint32(0x80000000) | uint32(ctrl.ControlType)<<16 | uint32(ctrl.Subtype)
	binary.BigEndian.PutUint32(buf[0:4], firstWord)
	binary.BigEndian.PutUint32(buf[4:8], ctrl.TypeSpecificInfo)
	binary.BigEndian.PutUint32(buf[8:12], ctrl.Timestamp)
	binary.BigEndian.PutUint32(buf[12:16], ctrl.DestinationSocketID)

	if len(ctrl.ControlInformationField) > 0 {
		copy(buf[headerSize:], ctrl.ControlInformationField)
	}

	return buf
}

// PeerIPFromUDPAddr converts a net.UDPAddr to the PeerIPAddress format.
// For IPv4, only IP1 is used (fields 2-4 are zero).
func PeerIPFromUDPAddr(addr *net.UDPAddr) PeerIPAddress {
	ip := addr.IP.To4()
	if ip != nil {
		// IPv4: stored in first 4 bytes, rest zero
		return PeerIPAddress{
			IP1: binary.BigEndian.Uint32(ip),
			IP2: 0,
			IP3: 0,
			IP4: 0,
		}
	}
	// IPv6: full 16 bytes
	ip6 := addr.IP.To16()
	return PeerIPAddress{
		IP1: binary.BigEndian.Uint32(ip6[0:4]),
		IP2: binary.BigEndian.Uint32(ip6[4:8]),
		IP3: binary.BigEndian.Uint32(ip6[8:12]),
		IP4: binary.BigEndian.Uint32(ip6[12:16]),
	}
}
