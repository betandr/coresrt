//                       Data Packet Structure

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |0|                    Packet Sequence Number                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |P P|O|K K|R|                   Message Number                  |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                              Data                             +
//    |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//    Packet Sequence Number: 31 bits.  The sequential number of the data
//       packet.

//    PP: 2 bits.  Packet Position Flag.  This field indicates the position
//       of the data packet in the message.  The value "10b" (binary) means
//       the first packet of the message. "00b" indicates a packet in the
//       middle. "01b" designates the last packet.  If a single data packet
//       forms the whole message, the value is "11b".

//    O: 1 bit.  Order Flag.  Indicates whether the message should be
//       delivered by the receiver in order (1) or not (0).  Certain
//       restrictions apply depending on the data transmission mode used
//       (Section 4.2).

//    KK: 2 bits.  Key-based Encryption Flag.  The flag bits indicate
//       whether or not data is encrypted.  The value "00b" (binary) means
//       data is not encrypted. "01b" indicates that data is encrypted with
//       an even key, and "10b" is used for odd key encryption.  Refer to
//       Section 6.  The value "11b" is only used in control packets.

//    R: 1 bit.  Retransmitted Packet Flag.  This flag is clear when a
//       packet is transmitted the first time.  The flag is set to "1" when
//       a packet is retransmitted.

//    Message Number: 26 bits.  The sequential number of consecutive data
//       packets that form a message (see PP field).

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

// Sharabayko, et al.        Expires 11 March 2022                 [Page 8]
//
// Internet-Draft                     SRT                    September 2021

// Data: variable length.  The payload of the data packet.  The length
//
//	of the data is the remaining length of the UDP packet.
package packets

type Data struct {
	PacketSequenceNumber    uint32
	PacketPositionFlag      byte   // 2 bits
	OrderFlag               byte   // 1 bit
	KeyBasedEncryptionFlag  byte   // 2 bits
	RetransmittedPacketFlag byte   // 1 bit
	MessageNumber           uint32 // 26 bits
	Timestamp               uint32
	DestinationSocketID     uint32
	Data                    []byte
}
