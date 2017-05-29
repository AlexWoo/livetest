package flv

import (
	"fmt"
	"io"
	"media"
)

var frameType = map[uint8]string{
	1: "keyframe",
	2: "interframe",
	3: "disposable-interframe",
	4: "generated-keyframe",
	5: "video-info/command-frame",
}

var codecID = map[uint8]string{
	1: "JPEG",
	2: "Sorenson-H.263",
	3: "Screen-Video",
	4: "VP6",
	5: "VP6-with-alpha-channel",
	6: "Screen-Video-V2",
	7: "AVC",
}

type FLVVideoData struct {
	FrameType   uint8
	CodecID     uint8
	VideoPacket media.MediaParser
	VideoData   []byte
}

func (data *FLVVideoData) Parser(r io.Reader, b []byte) (int, error) {
	var temp uint8

	temp = b[0]

	data.FrameType = temp & 0xf0 >> 4
	data.CodecID = temp & 0x0f
	data.VideoData = b[1:]

	switch data.CodecID {
	case 7:
		data.VideoPacket = new(AVCPacket)
		data.VideoPacket.Parser(r, data.VideoData)
	}

	return media.OK, nil
}

func (data *FLVVideoData) Sprint() string {
	switch data.CodecID {
	case 7:
		return fmt.Sprintf("[%s %s] %s",
			frameType[data.FrameType], codecID[data.CodecID],
			data.VideoPacket.Sprint())
	default:
		return fmt.Sprintf("[%s %s] Unsuppoted CodecID",
			frameType[data.FrameType], codecID[data.CodecID])

	}
}
