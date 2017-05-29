package flv

import (
	"fmt"
	"io"
	"media"
	"media/avc"
)

var avcPacketType = map[uint8]string{
	0: "AVC-SeqHeader",
	1: "AVC-Nalu",
	2: "AVC-endOfSeq",
}

type AVCPacket struct {
	AVCPacketType   uint8
	CompositionTime uint32
	Nalus           media.MediaParser
	AVCPayload      []byte
}

func (packet *AVCPacket) Parser(r io.Reader, b []byte) (int, error) {
	var err error

	packet.AVCPacketType = b[0]
	packet.CompositionTime = media.ReadUint24(b[1:])
	packet.AVCPayload = b[4:]

	switch packet.AVCPacketType {
	case 0:
		packet.Nalus = new(avc.AVCSeqHeader)
		_, err = packet.Nalus.Parser(r, packet.AVCPayload)
	case 1:
		packet.Nalus = new(avc.AVCNalus)
		_, err = packet.Nalus.Parser(r, packet.AVCPayload)
	default:
		return media.Error, fmt.Errorf("Unsupported AVCPacketType")
	}

	if err != nil {
		return media.Error, err
	}

	return media.OK, nil
}

func (packet *AVCPacket) Sprint() string {
	return fmt.Sprintf("%s cts=%d %s", avcPacketType[packet.AVCPacketType],
		packet.CompositionTime, packet.Nalus.Sprint())
}
