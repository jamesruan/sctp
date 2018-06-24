package format

import (
	"bytes"
	"encoding/binary"
)

type CommonHeader struct {
	SrcPort  uint16
	DstPort  uint16
	VTag     uint32
	CheckSum uint32 // CRC32c
}

type Packet struct {
	CommonHeader
	Chunks []ChunkField
}

func (s Packet) MarshalBinary() ([]byte, error) {
	var l uint16

	for _, v := range s.Chunks {
		l += v.Size()
	}

	buf := bytes.NewBuffer(make([]byte, 0, 12+l))
	binary.Write(buf, binary.BigEndian, s.CommonHeader)
	for _, v := range s.Chunks {
		v.WriteTo(buf)
	}
	return buf.Bytes(), nil
}
