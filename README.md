# Core Secure Reliable Transport Implementation in Go

_Basic implementation of the SRT protocol in Go._

[![Go](https://github.com/betandr/coresrt/actions/workflows/go.yml/badge.svg)](https://github.com/betandr/coresrt/actions/workflows/go.yml)

## Specification

[Network Working Group - The SRT Protocol](https://datatracker.ietf.org/doc/html/draft-sharabayko-srt-01)

## Send SRT

```
ffmpeg -i input.mp4 -c:v libx264 -b:v 3000k -c:a aac -b:a 128k \
  -f mpegts "srt://localhost:9999?pkt_size=1316&mode=caller&latency=2000"
```

### Parameters

- `pkt_size=1316`: Optimizes packet size for network transmission
- `mode=caller`: FFmpeg operates in caller mode (initiating the connection)
- `latency=2000`: Sets 2000ms latency buffer for error recovery

## Recieve SRT

```
go run .
```

## Receive SRT via ffmpeg for testing

```
ffmpeg -i "srt://localhost:9999?mode=listener" -c copy output.mkv
```

## Packet header

SRT packets are created at the application layer and handed to the transport layer for delivery. Each unit of SRT media or control data created by an application begins with the SRT packet header.

### SRT packet header

<table style="text-align:center">
    <tbody>
        <tr>
            <th><i>Offsets</i>
            </th>
            <th>Octet
            </th>
            <th colspan="8">0
            </th>
            <th colspan="8">1
            </th>
            <th colspan="8">2
            </th>
            <th colspan="8">3
            </th>
        </tr>
        <tr>
            <th>Octet
            </th>
            <th>Bit</th>
            <th>0
            </th>
            <th>1
            </th>
            <th colspan="1">2
            </th>
            <th colspan="1">3
            </th>
            <th>4
            </th>
            <th>5
            </th>
            <th>6
            </th>
            <th>7
            </th>
            <th colspan="1">8
            </th>
            <th>9
            </th>
            <th>10
            </th>
            <th>11
            </th>
            <th>12
            </th>
            <th>13
            </th>
            <th>14
            </th>
            <th>15
            </th>
            <th>16
            </th>
            <th>17
            </th>
            <th>18
            </th>
            <th>19
            </th>
            <th>20
            </th>
            <th>21
            </th>
            <th>22
            </th>
            <th>23
            </th>
            <th>24
            </th>
            <th>25
            </th>
            <th>26
            </th>
            <th>27
            </th>
            <th>28
            </th>
            <th>29
            </th>
            <th>30
            </th>
            <th>31
            </th>
        </tr>
        <tr align="center">
            <th>0
            </th>
            <th>0
            </th>
            <td colspan="1">F
            </td>
            <td colspan="31">Field meaning depends on the packet type
            </td>
        </tr>
        <tr align="center">
            <th>4
            </th>
            <th colspan="1">32
            </th>
            <td colspan="32">Field meaning depends on the packet type
            </td>
        </tr>
        <tr align="center">
            <th>8
            </th>
            <th>64
            </th>
            <td colspan="32">Timestamp
            </td>
        </tr>
        <tr align="center">
            <th>12
            </th>
            <th>96
            </th>
            <td colspan="32">Destination Socket ID
            </td>
        </tr>
        <tr align="center" style="height: 5em;">
            <th>...
            </th>
            <th>...
            </th>
            <td colspan="32">Packet Contents<br>(depends on the packet type)
            </td>
        </tr>
    </tbody>
</table>

### SRT data packet header

<table style="text-align:center">
    <tbody>
        <tr>
            <th><i>Offsets</i>
            </th>
            <th>Octet
            </th>
            <th colspan="8">0
            </th>
            <th colspan="8">1
            </th>
            <th colspan="8">2
            </th>
            <th colspan="8">3
            </th>
        </tr>
        <tr>
            <th>Octet
            </th>
            <th>Bit</th>
            <th>0
            </th>
            <th>1
            </th>
            <th colspan="1">2
            </th>
            <th colspan="1">3
            </th>
            <th>4
            </th>
            <th>5
            </th>
            <th>6
            </th>
            <th>7
            </th>
            <th colspan="1">8
            </th>
            <th>9
            </th>
            <th>10
            </th>
            <th>11
            </th>
            <th>12
            </th>
            <th>13
            </th>
            <th>14
            </th>
            <th>15
            </th>
            <th>16
            </th>
            <th>17
            </th>
            <th>18
            </th>
            <th>19
            </th>
            <th>20
            </th>
            <th>21
            </th>
            <th>22
            </th>
            <th>23
            </th>
            <th>24
            </th>
            <th>25
            </th>
            <th>26
            </th>
            <th>27
            </th>
            <th>28
            </th>
            <th>29
            </th>
            <th>30
            </th>
            <th>31
            </th>
        </tr>
        <tr align="center">
            <th>0
            </th>
            <th>0
            </th>
            <td colspan="1">0
            </td>
            <td colspan="31">Packet Sequence Number
            </td>
        </tr>
        <tr align="center">
            <th>4
            </th>
            <th>32
            </th>
            <td colspan="2">PP
            </td>
            <td>O
            </td>
            <td colspan="2">KK
            </td>
            <td>R
            </td>
            <td colspan="26">Message Number
            </td>
        </tr>
        <tr align="center">
            <th>8
            </th>
            <th>64
            </th>
            <td colspan="32">Timestamp
            </td>
        </tr>
        <tr align="center">
            <th>12
            </th>
            <th>96
            </th>
            <td colspan="32">Destination Socket ID
            </td>
        </tr>
        <tr align="center" style="height: 5em;">
            <th>...
            </th>
            <th>...
            </th>
            <td colspan="32">Data
            </td>
        </tr>
    </tbody>
</table>

The fields in the header are as follows:

- *Packet Sequence Number* (31 bits)
- *PP* (2 bits): Packet Position Flag
- *O* (1 bit): Order Flag
- *KK* (2 bits): Key-based Encryption Flag
- *R* (1 bit): Retransmitted Packet Flag
- *Message Number* (26 bits)
- *Data* (variable length)

### SRT control packet header

<table style="text-align:center">
    <tbody>
        <tr>
            <th><i>Offsets</i>
            </th>
            <th>Octet
            </th>
            <th colspan="8">0
            </th>
            <th colspan="8">1
            </th>
            <th colspan="8">2
            </th>
            <th colspan="8">3
            </th>
        </tr>
        <tr>
            <th>Octet
            </th>
            <th>Bit</th>
            <th>0
            </th>
            <th>1
            </th>
            <th colspan="1">2
            </th>
            <th colspan="1">3
            </th>
            <th>4
            </th>
            <th>5
            </th>
            <th>6
            </th>
            <th>7
            </th>
            <th colspan="1">8
            </th>
            <th>9
            </th>
            <th>10
            </th>
            <th>11
            </th>
            <th>12
            </th>
            <th>13
            </th>
            <th>14
            </th>
            <th>15
            </th>
            <th>16
            </th>
            <th>17
            </th>
            <th>18
            </th>
            <th>19
            </th>
            <th>20
            </th>
            <th>21
            </th>
            <th>22
            </th>
            <th>23
            </th>
            <th>24
            </th>
            <th>25
            </th>
            <th>26
            </th>
            <th>27
            </th>
            <th>28
            </th>
            <th>29
            </th>
            <th>30
            </th>
            <th>31
            </th>
        </tr>
        <tr align="center">
            <th>0
            </th>
            <th>0
            </th>
            <td colspan="1">1
            </td>
            <td colspan="15">Control Type
            </td>
            <td colspan="16">Subtype
            </td>
        </tr>
        <tr align="center">
            <th>4
            </th>
            <th>32
            </th>
            <td colspan="32">Type-specific Information
            </td>
        </tr>
        <tr align="center">
            <th>8
            </th>
            <th>64
            </th>
            <td colspan="32">Timestamp
            </td>
        </tr>
        <tr align="center">
            <th>12
            </th>
            <th>96
            </th>
            <td colspan="32">Destination Socket ID
            </td>
        </tr>
        <tr align="center" style="height: 5em;">
            <th>...
            </th>
            <th>...
            </th>
            <td colspan="32">Control Information Field (CIF)
            </td>
        </tr>
    </tbody>
</table>
