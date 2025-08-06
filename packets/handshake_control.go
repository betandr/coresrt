//    Handshake control packets (Control Type = 0x0000) are used to
//    exchange peer configurations, to agree on connection parameters, and
//    to establish a connection.

//    The Control Information Field (CIF) of a handshake control packet is
//    shown in Figure 5.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                            Version                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |        Encryption Field       |        Extension Field        |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                 Initial Packet Sequence Number                |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                 Maximum Transmission Unit Size                |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                    Maximum Flow Window Size                   |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         Handshake Type                        |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                         SRT Socket ID                         |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           SYN Cookie                          |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                                                               +
//    |                                                               |
//    +                        Peer IP Address                        +
//    |                                                               |
//    +                                                               +
//    |                                                               |
//    +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
//    |         Extension Type        |        Extension Length       |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                       Extension Contents                      +
//    |                                                               |
//    +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

//                     Figure 5: Handshake packet structure

//    Version: 32 bits.  A base protocol version number.  Currently used
//       values are 4 and 5.  Values greater than 5 are reserved for future
//       use.

//    Encryption Field: 16 bits.  Block cipher family and key size.  The
//       values of this field are described in Table 2.  The default value
//       is AES-128.

//                   +=======+============================+
//                   | Value | Cipher Family and Key Size |
//                   +=======+============================+
//                   | 0     |  No Encryption Advertised  |
//                   +-------+----------------------------+
//                   | 2     |          AES-128           |
//                   +-------+----------------------------+
//                   | 3     |          AES-192           |
//                   +-------+----------------------------+
//                   | 4     |          AES-256           |
//                   +-------+----------------------------+

//                       Table 2: Handshake Encryption
//                                Field values

//    Extension Field: 16 bits.  This field is message specific extension
//       related to Handshake Type field.  The value MUST be set to 0
//       except for the following cases.  (1) If the handshake control
//       packet is the INDUCTION message, this field is sent back by the
//       Listener. (2) In the case of a CONCLUSION message, this field
//       value should contain a combination of Extension Type values.  For
//       more details, see Section 4.3.1.

//                           +============+========+
//                           | Bitmask    |  Flag  |
//                           +============+========+
//                           | 0x00000001 | HSREQ  |
//                           +------------+--------+
//                           | 0x00000002 | KMREQ  |
//                           +------------+--------+
//                           | 0x00000004 | CONFIG |
//                           +------------+--------+

//                              Table 3: Handshake
//                               Extension Flags

//    Initial Packet Sequence Number: 32 bits.  The sequence number of the
//       very first data packet to be sent.

//    Maximum Transmission Unit Size: 32 bits.  This value is typically set
//       to 1500, which is the default Maximum Transmission Unit (MTU) size
//       for Ethernet, but can be less.

//    Maximum Flow Window Size: 32 bits.  The value of this field is the
//       maximum number of data packets allowed to be "in flight" (i.e. the
//       number of sent packets for which an ACK control packet has not yet
//       been received).

//    Handshake Type: 32 bits.  This field indicates the handshake packet
//       type.  The possible values are described in Table 4.  For more
//       details refer to Section 4.3.

//                       +============+================+
//                       | Value      | Handshake Type |
//                       +============+================+
//                       | 0xFFFFFFFD |      DONE      |
//                       +------------+----------------+
//                       | 0xFFFFFFFE |   AGREEMENT    |
//                       +------------+----------------+
//                       | 0xFFFFFFFF |   CONCLUSION   |
//                       +------------+----------------+
//                       | 0x00000000 |    WAVEHAND    |
//                       +------------+----------------+
//                       | 0x00000001 |   INDUCTION    |
//                       +------------+----------------+

//                           Table 4: Handshake Type

//    SRT Socket ID: 32 bits.  This field holds the ID of the source SRT
//       socket from which a handshake packet is issued.

//    SYN Cookie: 32 bits.  Randomized value for processing a handshake.
//       The value of this field is specified by the handshake message
//       type.  See Section 4.3.

//    Peer IP Address: 128 bits.  IPv4 or IPv6 address of the packet's
//       sender.  The value consists of four 32-bit fields.  In the case of
//       IPv4 addresses, fields 2, 3 and 4 are filled with zeroes.

//    Extension Type: 16 bits.  The value of this field is used to process
//       an integrated handshake.  Each extension can have a pair of
//       request and response types.

//             +=======+====================+===================+
//             | Value |   Extension Type   | HS Extension Flag |
//             +=======+====================+===================+
//             | 1     |   SRT_CMD_HSREQ    |       HSREQ       |
//             +-------+--------------------+-------------------+
//             | 2     |   SRT_CMD_HSRSP    |       HSREQ       |
//             +-------+--------------------+-------------------+
//             | 3     |   SRT_CMD_KMREQ    |       KMREQ       |
//             +-------+--------------------+-------------------+
//             | 4     |   SRT_CMD_KMRSP    |       KMREQ       |
//             +-------+--------------------+-------------------+
//             | 5     |    SRT_CMD_SID     |       CONFIG      |
//             +-------+--------------------+-------------------+
//             | 6     | SRT_CMD_CONGESTION |       CONFIG      |
//             +-------+--------------------+-------------------+
//             | 7     |   SRT_CMD_FILTER   |       CONFIG      |
//             +-------+--------------------+-------------------+
//             | 8     |   SRT_CMD_GROUP    |       CONFIG      |
//             +-------+--------------------+-------------------+

//                  Table 5: Handshake Extension Type values

//    Extension Length: 16 bits.  The length of the Extension Contents
//       field in four-byte blocks.

// Extension Contents: variable length.  The payload of the extension.
package packets

type CypherFamilyAndKeySize uint16

const (
	NoEncryption CypherFamilyAndKeySize = 0x0000
	AES128       CypherFamilyAndKeySize = 0x0002
	AES192       CypherFamilyAndKeySize = 0x0003
	AES256       CypherFamilyAndKeySize = 0x0004
)

type HandshakeExtensionFlag uint16

const (
	HSREQFlag  HandshakeExtensionFlag = 0x0001
	KMREQFlag  HandshakeExtensionFlag = 0x0002
	CONFIGFlag HandshakeExtensionFlag = 0x0004
)

type HandshakeType uint32

const (
	Done       HandshakeType = 0xFFFFFFFD
	Agreement  HandshakeType = 0xFFFFFFFE
	Conclusion HandshakeType = 0xFFFFFFFF
	WaveHand   HandshakeType = 0x00000000
	Induction  HandshakeType = 0x00000001
)

type ExtensionType uint16

const (
	HSREQ      ExtensionType = 1 // SRT_CMD_HSREQ
	HSRSP      ExtensionType = 2 // SRT_CMD_HSRSP
	KMREQ      ExtensionType = 3 // SRT_CMD_KMREQ
	KMRSP      ExtensionType = 4 // SRT_CMD_KMRSP
	SID        ExtensionType = 5 // SRT_CMD_SID
	Congestion ExtensionType = 6 // SRT_CMD_CONGESTION
	Filter     ExtensionType = 7 // SRT_CMD_FILTER
	Group      ExtensionType = 8 // SRT_CMD_GROUP
)

type PeerIPAddress struct {
	IP1 uint32
	IP2 uint32
	IP3 uint32
	IP4 uint32
}

type HandshakeControl struct {
	Version                     uint32                 // base protocol version number
	EncryptionField             CypherFamilyAndKeySize // Block cipher family and key size
	ExtensionField              HandshakeExtensionFlag // sequence number of the very first data packet to be sent.
	InitialPacketSequenceNumber uint32                 // default Maximum Transmission Unit (MTU) size for Ethernet, usually 1500
	MaximumTransmissionUnitSize uint32                 // the maximum number of data packets allowed to be "in flight"
	MaximumFlowWindowSize       uint32                 //
	HandshakeType               HandshakeType          // type of handshake packet
	SRTSocketID                 uint32                 // holds the ID of the source SRT socket from which a handshake packet is issued
	SYNCookie                   uint32                 // randomized value for processing a handshake
	PeerIPAddress               PeerIPAddress          // IPv4 or IPv6 address of the packet's sender
	ExtensionType               ExtensionType          // used to process an integrated handshake
	ExtensionLength             uint16
	ExtensionContents           []byte
}
