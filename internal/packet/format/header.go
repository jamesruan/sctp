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

func (s *CommonHeader) UnmarshalBinary(data []byte) (err error) {
	r := bytes.NewReader(data)
	err = binary.Read(r, binary.BigEndian, s)
	return
}

func (s Packet) MarshalBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, 1500))
	binary.Write(buf, binary.BigEndian, s.CommonHeader)
	for _, v := range s.Chunks {
		d, _ := v.MarshalBinary()
		buf.Write(d)
	}
	return buf.Bytes(), err
}
