//    3.2.10.  Peer Error

//    The Peer Error control packet is sent by a receiver when a processing
//    error (e.g. write to disk failure) occurs.  This informs the sender
//    of the situation and unblocks it from waiting for further responses
//    from the receiver.

//    The sender receiving this type of control packet must unblock any
//    sending operation in progress.

//    *NOTE*: This control packet is only used if the File Transfer
//    Congestion Control (Section 5.2) is enabled.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|      Control Type = 8       |         Reserved = 0          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                          Error Code                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                     Figure 19: Peer Error control packet

//    Packet Type: 1 bit, value = 1.  The packet type value of a Peer Error
//       control packet is "1".

//    Control Type: 15 bits, value = 8.  The control type value of a Peer
//       Error control packet is "8".

//    Error Code: 32 bits.  Peer error code.  At the moment the only value
//       defined is 4000 - file system error.

//    Timestamp: 32 bits.  See Section 3.

// Destination Socket ID: 32 bits.  See Section 3.
package packets

const (
	FileSystemErrorCode = 4000 // Error code for file system error
)

type PeerErrorControlPacket struct {
	ControlType uint16 // Value = 8. The control type value of a Peer Error control packet is "8".
	ErrorCode   uint32 // Error Code (currently only 4000 for file system error)
	Timestamp   uint32 // Timestamp
	DstSocketID uint32 // Destination Socket ID
}
