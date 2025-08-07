//    3.2.6.  Congestion Warning

//    The Congestion Warning control packet is reserved for future use.
//    Its purpose is to allow a receiver to signal a sender that there is
//    congestion happening at the receiving side.  The expected behaviour
//    is that upon receiving this packet the sender slows down its sending
//    rate by increasing the minimum inter-packet sending interval by a
//    discrete value (posited to be 12.5%).

//    Note that the conditions for a receiver to issue this type of packet
//    are not yet defined.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|      Control Type = 4       |         Reserved = 0          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                 Type-specific Information = 0                 |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                 Figure 15: Congestion Warning control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a Congestion
//       Warning control packet is "1".

//    Control Type: 15 bits, value = 4.  The control type value of a
//       Congestion Warning control packet is "4".

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

// Type-specific Information.  This field is reserved for future
//
//	definition.
package packets

type CongestionWarningControlPacket struct {
	ControlPacketType          // 1 bit, value = 1
	ControlType         uint16 // 15 bits, value = 4
	Timestamp           uint32 // 32 bits
	DestinationSocketID uint32 // 32 bits
}
