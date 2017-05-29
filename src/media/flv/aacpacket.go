package flv

import (
	"fmt"
	"io"
	"media"
)

var aacPacketType = map[uint8]string{
	0: "AAC-SeqHeader",
	1: "AAC-Raw",
}

type AACPacket struct {
	AACPacketType uint8
	AACPayload    []byte
}

func (packet *AACPacket) Parser(r io.Reader, b []byte) (int, error) {
	packet.AACPacketType = b[0]
	packet.AACPayload = b[1:]

	return media.OK, nil
}

func (packet *AACPacket) Sprint() string {
	return fmt.Sprintf("%s", aacPacketType[packet.AACPacketType])
}
