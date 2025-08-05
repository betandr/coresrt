//    3.2.3.  Keep-Alive

//    Keep-alive control packets are sent after a certain timeout from the
//    last time any packet (Control or Data) was sent.  The purpose of this
//    control packet is to notify the peer to keep the connection open when
//    no data exchange is taking place.

//    The default timeout for a keep-alive packet to be sent is 1 second.

//    An SRT keep-alive packet is formatted as follows:

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|         Control Type        |            Reserved           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                   Type-specific Information                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                     Figure 12: Keep-Alive control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a keep-alive
//       control packet is "1".

//    Control Type: 15 bits, value = KEEPALIVE{0x0001}.  The control type
//       value of a keep-alive control packet is "1".

//    Reserved: 16 bits, value = 0.  This is a fixed-width field reserved
//       for future use.

//    Type-specific Information.  This field is reserved for future
//       definition.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

// Keep-alive controls packet do not contain Control Information Field
// (CIF).
package packets

type KeepAliveControl struct {
	PacketType              uint8  // value = 1.  The packet type value of a keep-alive control packet is "1"
	ControlType             uint16 // 15 bits, value = KEEPALIVE{0x0001}.  The control type value of a keep-alive control packet is "1".
	Reserved                uint16 // 16 bits, value = 0.  This is a fixed-width field reserved for future use.
	TypeSpecificInformation uint32 // 32 bits.  This field is reserved for future definition.
	Timestamp               uint32 // 32 bits.  See Section 3.
	DestinationSocketID     uint32 // 32 bits.  See Section 3.
}
