//    3.2.9.  Message Drop Request

//    A Message Drop Request control packet is sent by the sender to the
//    receiver when a retransmission of an unacknowledged packet (forming a
//    whole or a part of a message) which is not present in the sender's
//    buffer is requested.  This may happen, for example, when a TTL
//    parameter (passed in the sending function) triggers a timeout for
//    retransmitting one or more lost packets which constitute parts of a
//    message, causing these packets to be removed from the sender's
//    buffer.

//    The sender notifies the receiver that it must not wait for
//    retransmission of this message.  Note that a Message Drop Request
//    control packet is not sent if the Too Late Packet Drop mechanism
//    (Section 4.6) causes the sender to drop a message, as in this case
//    the receiver is expected to drop it anyway.
//    A Message Drop Request contains the message number and corresponding
//    range of packet sequence numbers which form the whole message.  If
//    the sender does not already have in its buffer the specific packet or
//    packets for which retransmission was requested, then it is unable to
//    restore the message number.  In this case the Message Number field
//    must be set to zero, and the receiver should drop packets in the
//    provided packet sequence number range.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|      Control Type = 7       |         Reserved = 0          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                        Message Number                         |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                 First Packet Sequence Number                  |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                  Last Packet Sequence Number                  |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                    Figure 18: Drop Request control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a Drop
//       Request control packet is "1".

//    Control Type: 15 bits, value = 7.  The control type value of a Drop
//       Request control packet is "7".

//    Message Number: 32 bits.  The identifying number of the message
//       requested to be dropped.  See the Message Number field in
//       Section 3.1.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

//    First Packet Sequence Number: 32 bits.  The sequence number of the
//       first packet in the message.

// Last Packet Sequence Number: 32 bits.  The sequence number of the
//
//	last packet in the message.
package packets

type MessageDropRequest struct {
	ControlType         uint16 // Value = 7 The control type value of a Drop Request control packet is "7"
	MessageNumber       uint32 // dentifying number of the message requested to be dropped
	Timestamp           uint32
	DestinationSocketID uint32
	FirstPacketSeqNum   uint32 // Sequence number of the first packet in the message.
	LastPacketSeqNum    uint32 // Sequence number of the last packet in the message.
}
