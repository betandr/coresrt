package packets

// MinPacketSize is the minimum size of an SRT packet (16 bytes header).
const MinPacketSize = 16

// isControlPacket checks if the first bit of the packet
// is set, indicating a control packet.
func isControlPacket(data []byte) bool {
	return data[0]&0x80 != 0
}

// func ParsePacket(data []byte) (interface{}, error) {
// 	if len(data) < MinPacketSize {
// 		return nil, fmt.Errorf("packet too short: %d bytes (minimum %d)", len(data), MinPacketSize)
// 	}

// 	if isControlPacket(data) {
// 		return ParseControlPacket(data)
// 	}
// 	return ParseDataPacket(data)
// }
