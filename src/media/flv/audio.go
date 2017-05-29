package flv

import (
	"fmt"
	"io"
	"media"
)

var soundFormat = map[uint8]string{
	0:  "PCM-platform-endian",
	1:  "ADPCM",
	2:  "MP3",
	3:  "PCM-little-endian",
	4:  "Nellymoser-16kHz",
	5:  "Nellymoser-8kHz",
	6:  "Nellymoser",
	7:  "G.711-A-law",
	8:  "G.711-mu-law",
	9:  "reserved",
	10: "AAC",
	11: "Speex",
	14: "MP3-8kHz",
	15: "Device-specific",
}

var soundRate = map[uint8]string{
	0: "5.5kHz",
	1: "11kHz",
	2: "22kHz",
	3: "44kHz",
}

var soundSize = map[uint8]string{
	0: "8-Bit",
	1: "16-Bit",
}

var soundType = map[uint8]string{
	0: "Mono",
	1: "Stereo",
}

type FLVAudioData struct {
	SoundFormat uint8
	SoundRate   uint8
	SoundSize   uint8
	SoundType   uint8
	AudioPacket media.MediaParser
	SoundData   []byte
}

func (data *FLVAudioData) Parser(r io.Reader, b []byte) (int, error) {
	var temp uint8

	temp = b[0]

	data.SoundFormat = temp & 0xf0 >> 4
	data.SoundRate = temp & 0x0C >> 2
	data.SoundSize = temp & 0x02 >> 1
	data.SoundType = temp & 0x01
	data.SoundData = b[1:]

	switch data.SoundFormat {
	case 10:
		data.AudioPacket = new(AACPacket)
		data.AudioPacket.Parser(r, data.SoundData)
	}

	return media.OK, nil
}

func (data *FLVAudioData) Sprint() string {
	switch data.SoundFormat {
	case 10:
		return fmt.Sprintf("[%s %s %s %s] %s",
			soundFormat[data.SoundFormat], soundRate[data.SoundRate],
			soundSize[data.SoundSize], soundType[data.SoundType],
			data.AudioPacket.Sprint())
	default:
		return fmt.Sprintf("[%s %s %s %s] Unsuppoted SoundFormat",
			soundFormat[data.SoundFormat], soundRate[data.SoundRate],
			soundSize[data.SoundSize], soundType[data.SoundType])
	}
}
