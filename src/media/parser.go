package media

import (
	"bytes"
	"encoding/binary"
	"io"
)

const (
	OK    = 0
	Error = -1
	Again = -2
)

type MediaParser interface {
	Parser(r io.Reader, b []byte) (int, error)
	Sprint() string
}

func ReadUint16(b []byte) uint32 {
	var ui uint32
	var temp [4]byte

	temp[2] = b[0]
	temp[3] = b[1]
	sr := bytes.NewReader(temp[:])
	binary.Read(sr, binary.BigEndian, &ui)

	return ui
}

func ReadUint24(b []byte) uint32 {
	var ui uint32
	var temp [4]byte

	temp[1] = b[0]
	temp[2] = b[1]
	temp[3] = b[2]
	sr := bytes.NewReader(temp[:])
	binary.Read(sr, binary.BigEndian, &ui)

	return ui
}

func ReadUint32(b []byte) uint32 {
	var ui uint32

	sr := bytes.NewReader(b[0:4])
	binary.Read(sr, binary.BigEndian, &ui)

	return ui
}
