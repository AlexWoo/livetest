//Specification:
// ISO/IEC 14496-10 Advanced Video Coding
// ISO/IEC 14496-15 Advanced Video Coding (AVC) file format
package avc

import (
	"errors"
	"fmt"
	"io"
	"media"
)

var naluType = map[uint8]string{
	1:  "NALU_TYPE_SLICE",    //Coded slice of a non-IDR picture
	2:  "NALU_TYPE_DPA",      //Coded slice data partition A
	3:  "NALU_TYPE_DPB",      //Coded slice data partition B
	4:  "NALU_TYPE_DPC",      //Coded slice data partition C
	5:  "NALU_TYPE_IDR",      //Coded slice of an IDR picture
	6:  "NALU_TYPE_SEI",      //Supplemental enhancement information
	7:  "NALU_TYPE_SPS",      //Sequence parameter set
	8:  "NALU_TYPE_PPS",      //Picture parameter set
	9:  "NALU_TYPE_AUD",      //Access unit delimiter
	10: "NALU_TYPE_EOSEQ",    //End of sequence
	11: "NALU_TYPE_EOSTREAM", //End of stream
	12: "NALU_TYPE_FD",       //Filler data
	13: "NALU_TYPE_SPSE",     //Sequence parameter set extension
}

type AVCNalUnit struct {
	ForbidenZeroBit uint8
	NalRefIdc       uint8
	NaluType        uint8
}

func (nalu *AVCNalUnit) Parser(r io.Reader, b []byte) (int, error) {
	nalu.ForbidenZeroBit = b[0] & 0x80 >> 7
	nalu.NalRefIdc = b[0] & 0x60 >> 5
	nalu.NaluType = b[0] & 0x1F

	if nalu.ForbidenZeroBit != 0 {
		return media.Error, errors.New("forbidden_zero_bit != 0")
	}

	return media.OK, nil
}

func (nalu *AVCNalUnit) Sprint() string {
	return fmt.Sprintf("AVC nalutype=%s nri=%d",
		naluType[nalu.NaluType], nalu.NalRefIdc)
}
