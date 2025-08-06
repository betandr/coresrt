//                        SRT Packet Structure

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |F|        (Field meaning depends on the packet type)           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |          (Field meaning depends on the packet type)           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                        Packet Contents                        |
//    |                  (depends on the packet type)                 +
//    |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//    F: 1 bit.  Packet Type Flag.  The control packet has this flag set to
//       "1".  The data packet has this flag set to "0".

//    Timestamp: 32 bits.  The timestamp of the packet, in microseconds.
//       The value is relative to the time the SRT connection was
//       established.  Depending on the transmission mode (Section 4.2),
//       the field stores the packet send time or the packet origin time.

//    Destination Socket ID: 32 bits.  A fixed-width field providing the
//       SRT socket ID to which a packet should be dispatched.  The field
//       may have the special value "0" when the packet is a connection
//       request.

package packets

type Packet struct {
	PacketTypeFlag      byte // 1 bit for packet type (0 for data, 1 for control)
	Timestamp           uint32
	DestinationSocketID uint32
	Contents            []byte // Contents of the packet, type-specific
}
