//Specification:
// ISO/IEC 14496-10 Advanced Video Coding
// ISO/IEC 14496-15 Advanced Video Coding (AVC) file format
package avc

import (
	"fmt"
	"io"
	"media"
)

type AVCNalus struct {
	nalus []*AVCNalUnit
}

func (nalus *AVCNalus) Parser(r io.Reader, b []byte) (int, error) {
	var start uint32
	var end uint32

	end = uint32(len(b))
	i := 0

	for start < end {
		len := media.ReadUint32(b[start:])
		start += 4

		nalu := new(AVCNalUnit)
		_, err := nalu.Parser(r, b[start:start+len])
		if err != nil {
			return media.Error, err
		}

		start = start + len
		nalus.nalus = append(nalus.nalus, nalu)

		i += 1
	}

	if start == end {
		return media.OK, nil
	}

	return media.Error, fmt.Errorf("nalu[%d] decode error", i)
}

func (nalus *AVCNalus) Sprint() string {
	ret := "\n"

	for k, nalu := range nalus.nalus {
		ret = ret + fmt.Sprintf("\t%d %s\n", k, nalu.Sprint())
	}

	return ret
}
