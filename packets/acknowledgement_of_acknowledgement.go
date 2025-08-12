//    3.2.8.  ACKACK (Acknowledgement of Acknowledgement)

//    Acknowledgement of Acknowledgement (ACKACK) control packets are sent
//    to acknowledge the reception of a Full ACK and used in the
//    calculation of the round-trip time by the SRT receiver.

//    An SRT ACKACK control packet is formatted as follows:

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|        Control Type         |           Reserved            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Acknowledgement Number                    |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                       Figure 17: ACKACK control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of an ACKACK
//       control packet is "1".

//    Control Type: 15 bits, value = ACKACK{0x0006}.  The control type
//       value of an ACKACK control packet is "6".

//    Acknowledgement Number.  This field contains the Acknowledgement
//       Number of the full ACK packet the reception of which is being
//       acknowledged by this ACKACK packet.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

// ACKACK control packets do not contain Control Information Field
// (CIF).
package packets

type ACKACKControlPacket struct {
	PacketType            byte   // 1 bit, value = 1
	ControlType           uint16 // 15 bits, value = ACKACK{0x0006} value = 6
	AcknowledgementNumber uint32 // 32 bits
	Timestamp             uint32 // 32 bits
	DestinationSocketID   uint32 // 32 bits
}
