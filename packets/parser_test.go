package packets

import (
	"encoding/binary"
	"testing"
)

func TestIsControlPacket(t *testing.T) {
	control := []byte{0x80, 0x00, 0x00, 0x00}
	if !IsControlPacket(control) {
		t.Error("expected control packet (bit 0 set)")
	}

	data := []byte{0x00, 0x00, 0x00, 0x00}
	if IsControlPacket(data) {
		t.Error("expected data packet (bit 0 clear)")
	}
}

func TestParsePacketTooShort(t *testing.T) {
	_, err := ParsePacket([]byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("expected error for short packet")
	}
}

func TestParseDataPacket(t *testing.T) {
	// Build a data packet:
	// Byte 0-3: 0|seqnum=12345 (31 bits)
	// Byte 4-7: PP=11|O=1|KK=01|R=0|MsgNum=42 (26 bits)
	// Byte 8-11: Timestamp=1000000
	// Byte 12-15: DestSocketID=7
	// Byte 16+: payload "hello"
	raw := make([]byte, 21)

	// Sequence number: 12345 with bit 31 = 0
	binary.BigEndian.PutUint32(raw[0:4], 12345)

	// PP=11 (bits 31-30), O=1 (bit 29), KK=01 (bits 28-27), R=0 (bit 26), MsgNum=42 (bits 25-0)
	secondWord := uint32(0x03)<<30 | uint32(0x01)<<29 | uint32(0x01)<<27 | uint32(0x00)<<26 | uint32(42)
	binary.BigEndian.PutUint32(raw[4:8], secondWord)

	binary.BigEndian.PutUint32(raw[8:12], 1000000)
	binary.BigEndian.PutUint32(raw[12:16], 7)
	copy(raw[16:], []byte("hello"))

	pkt, err := ParseDataPacket(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pkt.PacketSequenceNumber != 12345 {
		t.Errorf("PacketSequenceNumber = %d, want 12345", pkt.PacketSequenceNumber)
	}
	if pkt.PacketPositionFlag != 3 {
		t.Errorf("PacketPositionFlag = %d, want 3", pkt.PacketPositionFlag)
	}
	if pkt.OrderFlag != 1 {
		t.Errorf("OrderFlag = %d, want 1", pkt.OrderFlag)
	}
	if pkt.KeyBasedEncryptionFlag != 1 {
		t.Errorf("KeyBasedEncryptionFlag = %d, want 1", pkt.KeyBasedEncryptionFlag)
	}
	if pkt.RetransmittedPacketFlag != 0 {
		t.Errorf("RetransmittedPacketFlag = %d, want 0", pkt.RetransmittedPacketFlag)
	}
	if pkt.MessageNumber != 42 {
		t.Errorf("MessageNumber = %d, want 42", pkt.MessageNumber)
	}
	if pkt.Timestamp != 1000000 {
		t.Errorf("Timestamp = %d, want 1000000", pkt.Timestamp)
	}
	if pkt.DestinationSocketID != 7 {
		t.Errorf("DestinationSocketID = %d, want 7", pkt.DestinationSocketID)
	}
	if string(pkt.Data) != "hello" {
		t.Errorf("Data = %q, want %q", string(pkt.Data), "hello")
	}
}

func TestParseControlPacket(t *testing.T) {
	// Build a control packet:
	// Byte 0-3: 1|ControlType=HANDSHAKE(0x0000)|Subtype=0x0000
	// Byte 4-7: TypeSpecificInfo=0
	// Byte 8-11: Timestamp=500
	// Byte 12-15: DestSocketID=0 (connection request)
	// Byte 16+: CIF data
	raw := make([]byte, 20)

	// First word: bit 31 set (control), control type = 0x0000, subtype = 0x0000
	firstWord := uint32(0x80000000)
	binary.BigEndian.PutUint32(raw[0:4], firstWord)

	binary.BigEndian.PutUint32(raw[4:8], 0)    // type-specific info
	binary.BigEndian.PutUint32(raw[8:12], 500)  // timestamp
	binary.BigEndian.PutUint32(raw[12:16], 0)   // dest socket ID
	copy(raw[16:], []byte{0xDE, 0xAD, 0xBE, 0xEF}) // CIF

	pkt, err := ParseControlPacket(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pkt.ControlType != HANDSHAKE {
		t.Errorf("ControlType = 0x%04X, want HANDSHAKE(0x0000)", uint16(pkt.ControlType))
	}
	if pkt.Subtype != 0 {
		t.Errorf("Subtype = 0x%04X, want 0x0000", uint16(pkt.Subtype))
	}
	if pkt.TypeSpecificInfo != 0 {
		t.Errorf("TypeSpecificInfo = %d, want 0", pkt.TypeSpecificInfo)
	}
	if pkt.Timestamp != 500 {
		t.Errorf("Timestamp = %d, want 500", pkt.Timestamp)
	}
	if pkt.DestinationSocketID != 0 {
		t.Errorf("DestinationSocketID = %d, want 0", pkt.DestinationSocketID)
	}
	if len(pkt.ControlInformationField) != 4 {
		t.Errorf("CIF length = %d, want 4", len(pkt.ControlInformationField))
	}
}

func TestParseControlPacketACK(t *testing.T) {
	raw := make([]byte, 16)
	// ACK control type = 0x0002
	firstWord := uint32(0x80000000) | uint32(0x0002)<<16
	binary.BigEndian.PutUint32(raw[0:4], firstWord)
	binary.BigEndian.PutUint32(raw[4:8], 1)     // ack number
	binary.BigEndian.PutUint32(raw[8:12], 9999) // timestamp
	binary.BigEndian.PutUint32(raw[12:16], 42)  // dest socket

	pkt, err := ParseControlPacket(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if pkt.ControlType != ACK {
		t.Errorf("ControlType = 0x%04X, want ACK(0x0002)", uint16(pkt.ControlType))
	}
}

func TestParsePacketDispatch(t *testing.T) {
	// Data packet (bit 0 clear)
	dataPkt := make([]byte, 16)
	binary.BigEndian.PutUint32(dataPkt[0:4], 1) // seq=1, bit 31=0

	result, err := ParsePacket(dataPkt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := result.(*Data); !ok {
		t.Errorf("expected *Data, got %T", result)
	}

	// Control packet (bit 0 set)
	ctrlPkt := make([]byte, 16)
	binary.BigEndian.PutUint32(ctrlPkt[0:4], 0x80050000) // SHUTDOWN

	result, err = ParsePacket(ctrlPkt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := result.(*Control); !ok {
		t.Errorf("expected *Control, got %T", result)
	}
}

func TestControlTypeName(t *testing.T) {
	tests := []struct {
		ct   ControlPacketType
		want string
	}{
		{HANDSHAKE, "HANDSHAKE"},
		{KEEPALIVE, "KEEPALIVE"},
		{ACK, "ACK"},
		{NAK, "NAK"},
		{CongestionWarning, "CONGESTION_WARNING"},
		{SHUTDOWN, "SHUTDOWN"},
		{ACKACK, "ACKACK"},
		{DROPREQ, "DROPREQ"},
		{PEERERROR, "PEERERROR"},
		{UserDefinedType, "USER_DEFINED"},
		{ControlPacketType(0x1234), "UNKNOWN(0x1234)"},
	}

	for _, tt := range tests {
		got := ControlTypeName(tt.ct)
		if got != tt.want {
			t.Errorf("ControlTypeName(%d) = %q, want %q", tt.ct, got, tt.want)
		}
	}
}
