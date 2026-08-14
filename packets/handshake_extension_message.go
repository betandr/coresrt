//    In a Handshake Extension, the value of the Extension Field of the
//    handshake control packet is defined as 1 for a Handshake Extension
//    request (SRT_CMD_HSREQ in Table 5), and 2 for a Handshake Extension
//    response (SRT_CMD_HSRSP in Table 5).

//    The Extension Contents field of a Handshake Extension Message is
//    structured as follows:

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                          SRT Version                          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           SRT Flags                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |      Receiver TSBPD Delay     |       Sender TSBPD Delay      |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//               Figure 6: Handshake Extension Message structure

//    SRT Version: 32 bits.  SRT library version MUST be formed as major *
//       0x10000 + minor * 0x100 + patch.

//    SRT Flags: 32 bits.  SRT configuration flags (see Section 3.2.1.1.1).

//    Receiver TSBPD Delay: 16 bits.  Timestamp-Based Packet Delivery
//       (TSBPD) Delay of the receiver.  Refer to Section 4.5.

//    Sender TSBPD Delay: 16 bits.  TSBPD of the sender.  Refer to
//       Section 4.5.

// 3.2.1.1.1.  Handshake Extension Message Flags

//                       +============+===============+
//                       | Bitmask    |      Flag     |
//                       +============+===============+
//                       | 0x00000001 |    TSBPDSND   |
//                       +------------+---------------+
//                       | 0x00000002 |    TSBPDRCV   |
//                       +------------+---------------+
//                       | 0x00000004 |     CRYPT     |
//                       +------------+---------------+
//                       | 0x00000008 |   TLPKTDROP   |
//                       +------------+---------------+
//                       | 0x00000010 |  PERIODICNAK  |
//                       +------------+---------------+
//                       | 0x00000020 |   REXMITFLG   |
//                       +------------+---------------+
//                       | 0x00000040 |     STREAM    |
//                       +------------+---------------+
//                       | 0x00000080 | PACKET_FILTER |
//                       +------------+---------------+

//                             Table 6: Handshake
//                          Extension Message Flags

//    *  TSBPDSND flag defines if the TSBPD mechanism (Section 4.5) will be
//       used for sending.

//    *  TSBPDRCV flag defines if the TSBPD mechanism (Section 4.5) will be
//       used for receiving.

//    *  CRYPT flag MUST be set.  It is a legacy flag that indicates the
//       party understands KK field of the SRT Packet (Figure 3).

//    *  TLPKTDROP flag should be set if too-late packet drop mechanism
//       will be used during transmission.  See Section 4.6.

//    *  PERIODICNAK flag set indicates the peer will send periodic NAK
//       packets.  See Section 4.8.2.

//    *  REXMITFLG flag MUST be set.  It is a legacy flag that indicates
//       the peer understands the R field of the SRT DATA Packet
//       (Figure 3).

//    *  STREAM flag identifies the transmission mode (Section 4.2) to be
//       used in the connection.  If the flag is set, the buffer mode
//       (Section 4.2.2) is used.  Otherwise, the message mode
//       (Section 4.2.1) is used.

// *  PACKET_FILTER flag indicates if the peer supports packet filter.
package packets

type HandshakeExtensionMessageFlags uint32

const (
	TSBPDSND     HandshakeExtensionMessageFlags = 0x00000001
	TSBPDRCV     HandshakeExtensionMessageFlags = 0x00000002
	CRYPT        HandshakeExtensionMessageFlags = 0x00000004
	TLPKTDROP    HandshakeExtensionMessageFlags = 0x00000008
	PERIODICNAK  HandshakeExtensionMessageFlags = 0x00000010
	REXMITFLG    HandshakeExtensionMessageFlags = 0x00000020
	STREAM       HandshakeExtensionMessageFlags = 0x00000040
	PACKETFILTER HandshakeExtensionMessageFlags = 0x00000080
)

type HandshakeExtensionMessage struct {
	SRTVersion         uint32                         // SRT library version formed as major * 0x10000 + minor * 0x100 + patch
	SRTFlags           HandshakeExtensionMessageFlags // SRT configuration flags
	ReceiverTSBPDDelay uint16                         // Timestamp-Based Packet Delivery (TSBPD) Delay of the receiver
	SenderTSBPDDelay   uint16                         // TSBPD of the sender
}
