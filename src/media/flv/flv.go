package flv

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"media"
)

type FileHeader struct {
	Signature            [3]byte
	Version              byte
	Flags                byte
	DataOffset           uint32
	FirstPreviousTagSize uint32
}

func (header *FileHeader) check() bool {
	if string(header.Signature[:]) != "FLV" {
		return false
	}

	if header.Version != 1 {
		return false
	}

	if header.DataOffset != 9 {
		return false
	}

	if header.FirstPreviousTagSize != 0 {
		return false
	}

	return true
}

type FLV struct {
	reader io.Reader
	buffer []byte
}

func Create(r io.Reader) *FLV {
	parser := new(FLV)
	parser.reader = r

	flvHeader := new(FileHeader)

	binary.Read(r, binary.BigEndian, flvHeader)
	if !flvHeader.check() {
		log.Fatalf("FLV Header(%v) Error, Not FLV file", flvHeader)
	}

	return parser
}

func (flv *FLV) Parser() {
	var seq uint32
	for {
		frame := new(FLVFrame)
		ret, err := frame.Parser(flv.reader, flv.buffer)
		switch ret {
		case media.Error:
			log.Fatal(err)
		case media.Again:
			continue
		default:
			frame.Seq = seq
			fmt.Println(frame.Sprint())
			seq += 1
		}
	}
}
