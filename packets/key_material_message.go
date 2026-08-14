//    3.2.2.  Key Material

//    The purpose of the Key Material Message is to let peers exchange
//    encryption-related information to be used to encrypt and decrypt the
//    payload of the stream.

//    This message can be supplied in two possible ways:

//    *  as a Handshake Extension (see Section 3.2.1.2)

//    *  in the Content Information Field of the User-Defined control
//       packet (described below).

//    When the Key Material is transmitted as a control packet, the Control
//    Type field of the SRT packet header is set to User-Defined Type (see
//    Table 1), the Subtype field of the header is set to SRT_CMD_KMREQ for
//    key-refresh request and SRT_CMD_KMRSP for key-refresh response
//    (Table 5).  The KM Refresh mechanism is described in Section 6.1.6.

//    The structure of the Key Material message is illustrated in
//    Figure 10.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |S|  V  |   PT  |              Sign             |   Resv1   | KK|
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                              KEKI                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |     Cipher    |      Auth     |       SE      |     Resv2     |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |             Resv3             |     SLen/4    |     KLen/4    |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                              Salt                             |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                          Wrapped Key                          +
//    |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                  Figure 10: Key Material Message structure

//    S: 1 bit, value = {0}.  This is a fixed-width field that is reserved
//       for future usage.

//    Version (V): 3 bits, value = {1}.  This is a fixed-width field that
//       indicates the SRT version:

//       *  1: Initial version.

//    Packet Type (PT): 4 bits, value = {2}.  This is a fixed-width field
//       that indicates the Packet Type:

//       *  0: Reserved

//       *  1: Media Stream Message (MSmsg)

//       *  2: Keying Material Message (KMmsg)

//       *  7: Reserved to discriminate MPEG-TS packet (0x47=sync byte).

//    Sign: 16 bits, value = {0x2029}.  This is a fixed-width field that
//       contains the signature 'HAI' encoded as a PnP Vendor ID [PNPID]
//       (in big-endian order).

//    Resv1: 6 bits, value = {0}.  This is a fixed-width field reserved for
//       flag extension or other usage.

//    Key-based Encryption (KK): 2 bits.  This is a fixed-width field that

//       indicates which SEKs (odd and/or even) are provided in the
//       extension:

//       *  00b: No SEK is provided (invalid extension format);

//       *  01b: Even key is provided;

//       *  10b: Odd key is provided;

//       *  11b: Both even and odd keys are provided.

//    Key Encryption Key Index (KEKI): 32 bits, value = {0}.  This is a
//       fixed-width field for specifying the KEK index (big-endian order)
//       was used to wrap (and optionally authenticate) the SEK(s).  The
//       value 0 is used to indicate the default key of the current stream.
//       Other values are reserved for the possible use of a key management
//       system in the future to retrieve a cryptographic context.

//       *  0: Default stream associated key (stream/system default)

//       *  1..255: Reserved for manually indexed keys.

//    Cipher: 8 bits, value = {0..2}.  This is a fixed-width field for
//       specifying encryption cipher and mode:

//       *  0: None or KEKI indexed crypto context

//       *  2: AES-CTR [SP800-38A].

//    Authentication (Auth): 8 bits, value = {0}.  This is a fixed-width
//       field for specifying a message authentication code algorithm:

//       *  0: None or KEKI indexed crypto context.

//    Stream Encapsulation (SE): 8 bits, value = {2}.  This is a fixed-
//       width field for describing the stream encapsulation:

//       *  0: Unspecified or KEKI indexed crypto context

//       *  1: MPEG-TS/UDP

//       *  2: MPEG-TS/SRT.

//    Resv2: 8 bits, value = {0}.  This is a fixed-width field reserved for
//       future use.

//    Resv3: 16 bits, value = {0}.  This is a fixed-width field reserved
//       for future use.

//    SLen/4: 8 bits, value = {4}.  This is a fixed-width field for
//       specifying salt length SLen in bytes divided by 4.  Can be zero if
//       no salt/IV present.  The only valid length of salt defined is 128
//       bits.

//    KLen/4: 8 bits, value = {4,6,8}.  This is a fixed-width field for
//       specifying SEK length in bytes divided by 4.  Size of one key even
//       if two keys present.  MUST match the key size specified in the
//       Encryption Field of the handshake packet Table 2.

//    Salt (SLen): SLen * 8 bits, value = { }.  This is a variable-width
//       field that complements the keying material by specifying a salt
//       key.

//    Wrap: (64 + n * KLen * 8) bits, value = { }.  This is a variable-
//       width field for specifying Wrapped key(s), where n = (KK + 1)/2
//       and the size of the wrap field is ((n * KLen) + 8) bytes.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    +                  Integrity Check Vector (ICV)                 +
//    |                                                               |
//    +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
//    |                              xSEK                             |
//    +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
//    |                              oSEK                             |
//    +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

//                      Figure 11: Unwrapped key structure

//    ICV: 64 bits.  64-bit Integrity Check Vector(AES key wrap integrity).
//       This field is used to detect if the keys were unwrapped properly.
//       If the KEK in hand is invalid, validation fails and unwrapped keys
//       are discarded.

//    xSEK: variable width.  This field identifies an odd or even SEK.  If
//       only one key is present, the bit set in the KK field tells which
//       SEK is provided.  If both keys are present, then this field is
//       eSEK (even key) and it is followed by odd key oSEK.  The length of
//       this field is calculated as KLen * 8.

//    oSEK: variable width.  This field with the odd key is present only
//       when the message carries the two SEKs (identified by he KK field).

package packets

type KeyMaterialPacketType byte

const (
	Reserved       KeyMaterialPacketType = 0
	MediaStream    KeyMaterialPacketType = 1 // Media Stream Message (MSmsg)
	KeyingMaterial KeyMaterialPacketType = 2 // Keying Material Message (KMmsg)
	MPEGTSReserved KeyMaterialPacketType = 7 // Reserved to discriminate MPEG-TS packet (0x47=sync byte)
)

type KeyBasedEncryption byte

const (
	NoSEKProvided KeyBasedEncryption = 0b00 // No SEK is provided (invalid extension format)
	EvenKey       KeyBasedEncryption = 0b01 // Even key is provided
	OddKey        KeyBasedEncryption = 0b10 // Odd key is provided
	BothKeys      KeyBasedEncryption = 0b11 // Both even and odd keys are provided
)

type KeyMaterialCipher byte

const (
	NoneOrKEKI KeyMaterialCipher = 0 // None or KEKI indexed crypto context
	AESCTR     KeyMaterialCipher = 2 // AES-CTR [SP800-38A]
)

type KeyMaterialAuthentication byte

const (
	NoAuthOrKEKI KeyMaterialAuthentication = 0 // None or KEKI indexed crypto context
)

type KeyMaterialStreamEncapsulation byte

const (
	UnspecifiedOrKEKI KeyMaterialStreamEncapsulation = 0 // Unspecified or KEKI indexed crypto context
	MPEGTSUDP         KeyMaterialStreamEncapsulation = 1 // MPEG-TS/UDP
	MPEGTSSRT         KeyMaterialStreamEncapsulation = 2 // MPEG-TS/SRT
)

type KeyMaterialMessage struct {
	S                     byte                           // 1 bit, reserved for future usage
	Version               byte                           // 3 bits, SRT version
	PacketType            KeyMaterialPacketType          // 4 bits, Packet Type
	Sign                  uint16                         // 16 bits, value = {0x2029} signature 'HAI' encoded as a PnP Vendor ID
	Resv1                 uint8                          // 6 bits, value = {0}.  This is a fixed-width field reserved for future use.
	KeyBasedEncryption    byte                           // 2 bits, indicates which SEKs (odd and/or even) are provided
	KeyEncryptionKeyIndex uint32                         // fixed-width field for specifying the KEK index (big-endian order)was used to wrap (and optionally authenticate) the SEK(s)
	Cipher                KeyMaterialCipher              // 8 bits, specifies encryption cipher and mode
	Authentication        KeyMaterialAuthentication      // 8 bits, specifies a message authentication code algorithm
	StreamEncapsulation   KeyMaterialStreamEncapsulation // 8 bits, describes the stream encapsulation
	Resv2                 uint8                          // 8 bits, reserved for future use
	Resv3                 uint16                         // 16 bits, reserved for future use
	SaltLength            uint8                          // 8 bits, value = {4} specifies salt length SLen in bytes divided by 4
	KeyLength             uint8                          // 8 bits, value = {4,6,8} specifies SEK length in bytes divided by 4
	Salt                  []byte                         // variable-width field that complements the keying material by
	Wrap                  []byte                         // variable-width field for specifying Wrapped key(s)
	IntegrityCheckVector  []byte                         // 64 bits, Integrity Check Vector (ICV) for AES key wrap integrity
	Xsek                  []byte                         // variable width, identifies an odd or even SEK
	Osek                  []byte                         // variable width, identifies an odd SEK
}
