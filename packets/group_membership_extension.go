//    3.2.1.4.  Group Membership Extension

//    The Group Membership handshake extension is reserved for the future
//    and is going to be used to allow multipath SRT connections.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                           Group ID                            |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |     Type    |     Flags     |             Weight              |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                 Figure 8: Group Membership Extension Message

//    GroupID: 32 bits.  The identifier of a group whose members include
//       the sender socket that is making a connection.  The target socket
//       that is interpreting GroupID SHOULD belong to the corresponding
//       group on the target side.  If such a group does not exist, the
//       target socket MAY create it.

//    Type: 8 bits.  Group type, as per SRT_GTYPE_ enumeration:

//    *  0: undefined group type,

//    *  1: broadcast group type,

//    *  2: main/backup group type,

//    *  3: balancing group type,

//    *  4: multicast group type (reserved for future use).

//    Flags: 8 bits.  Special flags mostly reserved for the future.  See
//       Figure 9.

//    Weight: 16 bits.  Special value with interpretation depending on the
//       Type field value:

//    *  Not used with broadcast group type,

//    *  Defines the link priority for main/backup group type,

//    *  Not yet defined for any other cases (reserved for future use).

//     0 1 2 3 4 5 6 7
//    +-+-+-+-+-+-+-+
//    |   (zero)  |M|
//    +-+-+-+-+-+-+-+

//                  Figure 9: Group Membership Extension Flags

// M: 1 bit.  When set, defines synchronization on message numbers,
//
//	otherwise transmission is synchronized on sequence numbers.
package packets

type SrtGtype uint8

const (
	GTYPE_UNDEFINED   SrtGtype = 0 // undefined group type
	GTYPE_BROADCAST   SrtGtype = 1 // broadcast group type
	GTYPE_MAIN_BACKUP SrtGtype = 2 // main/backup group type
	GTYPE_BALANCING   SrtGtype = 3 // balancing group type
	GTYPE_MULTICAST   SrtGtype = 4 // Reserved for future use
)

type GroupMembershipExtension struct {
	GroupID uint32   // 32 bits for group ID
	Type    SrtGtype // 8 bits for group type
	Flags   uint8    // 8 bits for special flags
	Weight  uint16   // 16 bits for link priority
}
