//Specification:
// ISO_IEC 14496-10 Advanced Video Coding
// ISO_IEC 14496-15 Advanced Video Coding (AVC) file format
package avc

import (
	"fmt"
	"io"
	"media"
)

type AVCSeqHeader struct {
	ConfigurationVersion  uint8
	AVCProfileIndication  uint8
	ProfileCompatibility  uint8
	AVCLevelIndication    uint8
	CompleteRepresenation uint8
	LengthSizeMinusOne    uint8
	NrOfSPS               uint8
	SPSs                  []*AVCNalUnit
	NrOfPPS               uint8
	PPSs                  []*AVCNalUnit
}

// ISO/IEC 14496-15 Advanced Video Coding (AVC) file format 5.2.4.1.1
func (seqHeader *AVCSeqHeader) Parser(r io.Reader, b []byte) (int, error) {
	var start uint32
	var end uint32

	start = 6
	end = uint32(len(b))

	seqHeader.ConfigurationVersion = b[0]
	seqHeader.AVCProfileIndication = b[1]
	seqHeader.ProfileCompatibility = b[2]
	seqHeader.AVCLevelIndication = b[3]
	seqHeader.CompleteRepresenation = b[4] & 0x80 >> 7
	seqHeader.LengthSizeMinusOne = b[4] & 0x03

	seqHeader.NrOfSPS = b[5] & 0x1F
	for i := 0; i < int(seqHeader.NrOfSPS); i += 1 {
		len := media.ReadUint16(b[start:])
		start += 2

		nalu := new(AVCNalUnit)
		_, err := nalu.Parser(r, b[start:start+len])
		if err != nil {
			return media.Error, fmt.Errorf("SPS[%d] decode error", i)
		}
		seqHeader.SPSs = append(seqHeader.SPSs, nalu)
		start = start + len
	}

	seqHeader.NrOfPPS = b[start]
	start += 1
	for i := 0; i < int(seqHeader.NrOfPPS); i += 1 {
		len := media.ReadUint16(b[start:])
		start += 2

		nalu := new(AVCNalUnit)
		_, err := nalu.Parser(r, b[start:start+len])
		if err != nil {
			return media.Error, fmt.Errorf("PPS[%d] decode error", i)
		}
		seqHeader.PPSs = append(seqHeader.PPSs, nalu)
		start = start + len
	}

	if start > end {
		return media.Error, fmt.Errorf("Sequnece Header decode error")
	}

	return media.OK, nil
}

func (seqHeader *AVCSeqHeader) Sprint() string {
	var ret string

	ret = fmt.Sprintf("confVer=%d AVC_pid=%d pc=%d AVC_LI=%d CR=%d LSMO=%d\n",
		seqHeader.ConfigurationVersion, seqHeader.AVCProfileIndication,
		seqHeader.ProfileCompatibility, seqHeader.AVCLevelIndication,
		seqHeader.CompleteRepresenation, seqHeader.LengthSizeMinusOne)

	for k, nalu := range seqHeader.SPSs {
		ret = ret + fmt.Sprintf("\tSPS[%d] %s\n", k, nalu.Sprint())
	}

	for k, nalu := range seqHeader.PPSs {
		ret = ret + fmt.Sprintf("\tPPS[%d] %s\n", k, nalu.Sprint())
	}

	return ret
}
