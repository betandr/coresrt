//                      Control Packet Structure

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+- SRT Header +-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |1|         Control Type        |            Subtype            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                   Type-specific Information                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Timestamp                           |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                     Destination Socket ID                     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+- CIF -+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                   Control Information Field                   +
//    |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//    Control Type: 15 bits.  Control Packet Type.  The use of these bits
//       is determined by the control packet type definition.  See Table 1.

//    Subtype: 16 bits.  This field specifies an additional subtype for
//       specific packets.  See Table 1.

//    Type-specific Information: 32 bits.  The use of this field depends on
//       the particular control packet type.  Handshake packets do not use
//       this field.

//    Timestamp: 32 bits.  See Section 3.

//    Destination Socket ID: 32 bits.  See Section 3.

//    Control Information Field (CIF): variable length.  The use of this
//       field is defined by the Control Type field of the control packet.

//    The types of SRT control packets are shown in Table 1.  The value
//    "0x7FFF" is reserved for a user-defined type.

//      +====================+==============+=========+================+
//      | Packet Type        | Control Type | Subtype | Section        |
//      +====================+==============+=========+================+
//      | HANDSHAKE          |    0x0000    |   0x0   | Section 3.2.1  |
//      +--------------------+--------------+---------+----------------+
//      | KEEPALIVE          |    0x0001    |   0x0   | Section 3.2.3  |
//      +--------------------+--------------+---------+----------------+
//      | ACK                |    0x0002    |   0x0   | Section 3.2.4  |
//      +--------------------+--------------+---------+----------------+
//      | NAK (Loss Report)  |    0x0003    |   0x0   | Section 3.2.5  |
//      +--------------------+--------------+---------+----------------+
//      | Congestion Warning |    0x0004    |   0x0   | Section 3.2.6  |
//      +--------------------+--------------+---------+----------------+
//      | SHUTDOWN           |    0x0005    |   0x0   | Section 3.2.7  |
//      +--------------------+--------------+---------+----------------+
//      | ACKACK             |    0x0006    |   0x0   | Section 3.2.8  |
//      +--------------------+--------------+---------+----------------+
//      | DROPREQ            |    0x0007    |   0x0   | Section 3.2.9  |
//      +--------------------+--------------+---------+----------------+
//      | PEERERROR          |    0x0008    |   0x0   | Section 3.2.10 |
//      +--------------------+--------------+---------+----------------+
//      | User-Defined Type  |    0x7FFF    |    -    | N/A            |
//      +--------------------+--------------+---------+----------------+

// Table 1: SRT control packet types
package packets

type ControlPacketType uint16

const (
	HANDSHAKE         ControlPacketType = 0x0000
	KEEPALIVE         ControlPacketType = 0x0001
	ACK               ControlPacketType = 0x0002
	NAK               ControlPacketType = 0x0003 // Loss report
	CongestionWarning ControlPacketType = 0x0004
	SHUTDOWN          ControlPacketType = 0x0005
	ACKACK            ControlPacketType = 0x0006
	DROPREQ           ControlPacketType = 0x0007
	PEERERROR         ControlPacketType = 0x0008
	UserDefinedType   ControlPacketType = 0x7FFF // Reserved for a user-defined type
)

type Control struct {
	ControlType             ControlPacketType // 15 bits for control type
	Subtype                 ControlPacketType // 16 bits for subtype
	TypeSpecificInfo        uint32            // 32 bits for type-specific information
	Timestamp               uint32
	DestinationSocketID     uint32
	ControlInformationField []byte // Control Information Field
}
