package flv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"media"
)

var tagType = map[uint8]string{
	8:  "audio",
	9:  "video",
	18: "script",
}

type FLVTag struct {
	TagType   uint8
	DataSize  uint32
	Timestamp uint32
	StreamID  uint32
}

func readFLVTimestamp(b []byte) uint32 {
	var ui uint32
	var temp [4]byte

	temp[0] = b[3]
	temp[1] = b[0]
	temp[2] = b[1]
	temp[3] = b[2]
	sr := bytes.NewReader(temp[:])
	binary.Read(sr, binary.BigEndian, &ui)

	return ui
}

func (tag *FLVTag) Parser(r io.Reader, b []byte) (int, error) {
	flvtagForParser := make([]byte, 11)
	binary.Read(r, binary.BigEndian, flvtagForParser)

	tag.TagType = flvtagForParser[0]
	tag.DataSize = media.ReadUint24(flvtagForParser[1:])
	tag.Timestamp = readFLVTimestamp(flvtagForParser[4:])
	tag.StreamID = media.ReadUint24(flvtagForParser[8:])

	return media.OK, nil
}

func (tag *FLVTag) Sprint() string {
	return fmt.Sprintf("type:%s mlen:%d timestamp:%d sid:%d",
		tagType[tag.TagType], tag.DataSize, tag.Timestamp, tag.StreamID)
}
