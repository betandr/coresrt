//    3.2.1.3.  Stream ID Extension Message

//    The Stream ID handshake extension message can be used to identify the
//    stream content.  The Stream ID value can be free-form, but there is
//    also a recommended convention that can be used to achieve
//    interoperability.

//    The Stream ID handshake extension message has SRT_CMD_SID extension
//    type (see Table 5.  The extension contents are a sequence of UTF-8
//    characters.  The maximum allowed size of the StreamID extension is
//    512 bytes.

//     0                   1                   2                   3
//     0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//    |                                                               |
//    |                           Stream ID                           |
//                                   ...
//    |                                                               |
//    +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

//                    Figure 7: Stream ID Extension Message

//    The Extension Contents field holds a sequence of UTF-8 characters
//    (see Figure 7).  The maximum allowed size of the StreamID extension
//    is 512 bytes.  The actual size is determined by the Extension Length
//    field (Figure 5), which defines the length in four byte blocks.  If
//    the actual payload is less than the declared length, the remaining
//    bytes are set to zeros.

//    The content is stored as 32-bit little endian words.

package packets

type StreamIdExtensionMessage struct {
	StreamID string
}
