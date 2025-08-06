//    3.2.5.  NAK (Negative Acknowledgement or Loss Report)

//    Negative acknowledgment (NAK) control packets are used to signal
//    failed data packet deliveries.  The receiver notifies the sender
//    about lost data packets by sending a NAK packet that contains a list
//    of sequence numbers for those lost packets.

//    An SRT NAK packet is formatted as follows:

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
//    +-+-+-+-+-+-+-+-+-+-+-+- CIF (Loss List) -+-+-+-+-+-+-+-+-+-+-+-+
//    |0|                 Lost packet sequence number                 |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|         Range of lost packets from sequence number          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |0|                    Up to sequence number                    |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |0|                 Lost packet sequence number                 |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                        Figure 14: NAK control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a NAK
//       control packet is "1".

//    Control Type: 15 bits, value = NAK{0x0003}.  The control type value
//       of a NAK control packet is "3".

//    Reserved: 16 bits, value = 0.  This is a fixed-width field reserved
//       for future use.

//    Type-specific Information: 32 bits.  This field is reserved for
//       future definition.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

// Control Information Field (CIF).  A single value or a range of lost
//
//	packets sequence numbers.  See packet sequence number coding in
//	Appendix A.
package packets

type NegativeAcknowledgmentControlPacket struct {
	PacketType              uint8    // value = 1.  The packet type value of a NAK control packet is "1"
	ControlType             uint16   // value = NAK{0x0003}.  The control type value of a NAK control packet is "3"
	Reserved                uint32   // value = 0.  This is a fixed-width field reserved for future use.
	TypeSpecificInformation uint32   // reserved for future definition
	Timestamp               uint32   // See Section 3.
	DestinationSocketID     uint32   // See Section 3.
	ControlInformationField []uint32 // Control Information Field (CIF).  A single value or a range of lost packets sequence numbers.
}
