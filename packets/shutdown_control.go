//    3.2.7.  Shutdown

//    Shutdown control packets are used to initiate the closing of an SRT
//    connection.

//    An SRT shutdown control packet is formatted as follows:

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|        Control Type         |           Reserved            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                   Type-specific Information                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                      Figure 16: Shutdown control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a shutdown
//       control packet is "1".

//    Control Type: 15 bits, value = SHUTDOWN{0x0005}.  The control type
//       value of a shutdown control packet is "5".

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

//    Type-specific Information.  This field is reserved for future
//       definition.

// Shutdown control packets do not contain Control Information Field
// (CIF).
package packets

type ShutdownControlPacket struct {
	ControlPacketType          // 1 bit, value = 1
	ControlType         uint16 // 15 bits, value = SHUTDOWN{0x0005} value = 5
	Timestamp           uint32 // 32 bits
	DestinationSocketID uint32 // 32 bits
}
