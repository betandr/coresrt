//    3.2.4.  ACK (Acknowledgment)

//    Acknowledgment (ACK) control packets are used to provide the delivery
//    status of data packets.  By acknowledging the reception of data
//    packets up to the acknowledged packet sequence number, the receiver
//    notifies the sender that all prior packets were received or, in the
//    case of live streaming (Section 4.2, Section 7.1), preceding missing
//    packets (if any) were dropped as too late to be delivered
//    (Section 4.6).

//    ACK packets may also carry some additional information from the
//    receiver like the estimates of RTT, RTT variance, link capacity,
//    receiving speed, etc.  The CIF portion of the ACK control packet is
//    expanded as follows:

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|        Control Type         |           Reserved            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                    Acknowledgement Number                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+- CIF -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |            Last Acknowledged Packet Sequence Number           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                              RTT                              |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                          RTT Variance                         |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Available Buffer Size                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Packets Receiving Rate                    |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Estimated Link Capacity                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         Receiving Rate                        |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                        Figure 13: ACK control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of an ACK
//       control packet is "1".

//    Control Type: 15 bits, value = ACK{0x0002}.  The control type value
//       of an ACK control packet is "2".

//    Reserved: 16 bits, value = 0.  This is a fixed-width field reserved
//       for future use.

//    Acknowledgement Number: 32 bits.  This field contains the sequential
//       number of the full acknowledgment packet starting from 1.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

//    Last Acknowledged Packet Sequence Number: 32 bits.  This field
//       contains the sequence number of the last data packet being
//       acknowledged plus one.  In other words, if it the sequence number
//       of the first unacknowledged packet.

//    RTT: 32 bits.  RTT value, in microseconds, estimated by the receiver
//       based on the previous ACK/ACKACK packet pair exchange.

//    RTT Variance: 32 bits.  The variance of the RTT estimate, in
//       microseconds.

//    Available Buffer Size: 32 bits.  Available size of the receiver's
//       buffer, in packets.

//    Packets Receiving Rate: 32 bits.  The rate at which packets are being
//       received, in packets per second.

//    Estimated Link Capacity: 32 bits.  Estimated bandwidth of the link,
//       in packets per second.

//    Receiving Rate: 32 bits.  Estimated receiving rate, in bytes per
//       second.

//    There are several types of ACK packets:

//    *  A Full ACK control packet is sent every 10 ms and has all the
//       fields of Figure 13.

//    *  A Light ACK control packet includes only the Last Acknowledged
//       Packet Sequence Number field.  The Type-specific Information field
//       should be set to 0.

//    *  A Small ACK includes the fields up to and including the Available
//       Buffer Size field.  The Type-specific Information field should be
//       set to 0.

//    The sender only acknowledges the receipt of Full ACK packets (see
//    Section 3.2.8).

//    The Light ACK and Small ACK packets are used in cases when the
//    receiver should acknowledge received data packets more often than
//    every 10 ms.  This is usually needed at high data rates.  It is up to
//    the receiver to decide the condition and the type of ACK packet to
//    send (Light or Small).  The recommendation is to send a Light ACK for
//    every 64 packets received.

// 3.2.5.  NAK (Negative Acknowledgement or Loss Report)

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

type AcknowledgementControlPacket struct {
	PacketType                           uint8  // value = 1.  The packet type value of an acknowledgment control packet is "1"
	ControlType                          uint16 // 15 bits, value = ACK{0x0002}.  The control type value of an acknowledgment control packet is "2".
	Reserved                             uint16 // 16 bits, value = 0.  This is a fixed-width field reserved for future use.
	AcknowledgementNumber                uint32 // 32 bits.  The sequential number of the full acknowledgment packet starting from 1.
	Timestamp                            uint32 // 32 bits.
	DestinationSocketID                  uint32 // 32 bits.
	LastAcknowledgedPacketSequenceNumber uint32 // 32 bits.  The sequence number of the last data packet being acknowledged plus one.
	RTT                                  uint32 // 32 bits.  RTT value, in microseconds, estimated by the receiver based on the previous ACK/ACKACK packet pair exchange.
	RTTVariance                          uint32 // 32 bits.  The variance of the RTT estimate, in microseconds.
	AvailableBufferSize                  uint32 // 32 bits.  Available size of the receiver's buffer, in packets.
	PacketsReceivingRate                 uint32 // 32 bits.
	EstimatedLinkCapacity                uint32 // 32 bits.  Estimated bandwidth of the link,
	ReceivingRate                        uint32 // 32 bits.  Estimated receiving rate, in bytes per
}
