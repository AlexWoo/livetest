package flv

import (
	"errors"
	"fmt"
	"io"
	"media"
)

const (
	FLVAudio  = 8
	FLVVideo  = 9
	FLVScript = 18
)

type FLVFrame struct {
	flvTag  *FLVTag
	flvData media.MediaParser
	Seq     uint32
	n       uint32
	buffer  []byte
}

func (frame *FLVFrame) Parser(r io.Reader, b []byte) (int, error) {
	var err error

	// Parser Tag
	if frame.n == 0 {
		frame.flvTag = new(FLVTag)
		frame.flvTag.Parser(r, b)

		frame.buffer = make([]byte, frame.flvTag.DataSize)
	}

	// Parser Data
	n, err := r.Read(frame.buffer[frame.n:])
	if err != nil {
		return media.Error, err
	}

	frame.n += uint32(n)
	if frame.n < frame.flvTag.DataSize {
		return media.Again, nil
	}

	switch frame.flvTag.TagType {
	case FLVAudio:
		frame.flvData = new(FLVAudioData)
		_, err = frame.flvData.Parser(r, frame.buffer)
	case FLVVideo:
		frame.flvData = new(FLVVideoData)
		_, err = frame.flvData.Parser(r, frame.buffer)
	case FLVScript:
		frame.flvData = new(FLVScriptData)
		_, err = frame.flvData.Parser(r, frame.buffer)
	default:
		return media.Error, errors.New("Unknown TagType")
	}

	if err != nil {
		return media.Error, err
	}

	// Parser PreviousTagSize
	var temp [4]byte
	n, err = r.Read(temp[:])
	if err != nil || n != 4 {
		return media.Error, err
	}
	previousTagSize := media.ReadUint32(temp[:])
	if previousTagSize != frame.flvTag.DataSize+11 {
		return media.Error, errors.New("previousTagSize error")
	}

	return media.OK, nil
}

func (frame *FLVFrame) Sprint() string {
	return fmt.Sprintf("%d flvframe %s %s", frame.Seq, frame.flvTag.Sprint(),
		frame.flvData.Sprint())
}
