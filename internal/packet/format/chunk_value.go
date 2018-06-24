package format

import (
	"encoding/binary"
	"io"
)

type ChunkDATA struct {
	ChunkFieldHeader
	ChunkDATAValue
}

func (s ChunkDATA) ToChunkField() ChunkField {
	return ChunkField{s.ChunkFieldHeader, s.ChunkDATAValue}
}

func (s *ChunkDATA) SetData(data []byte) {
	s.Data = data
	s.ChunkFieldHeader.Length = s.ChunkDATAValue.Len()
	padlen := s.ChunkDATAValue.Size() - s.ChunkDATAValue.Len()
	s.Padding = padding4[:padlen]
}

type ChunkDATAValue struct {
	ChunkDATAValueHeader
	Data    []byte
	Padding []byte
}

type ChunkDATAValueHeader struct {
	TSN      uint32
	StreamID uint16
	StreamSN uint16
	PPI      uint32 // payload  protocol identifier
}

func (v ChunkDATAValue) Len() uint16 {
	return uint16(len(v.Data) + 12)
}

func (v ChunkDATAValue) Size() uint16 {
	return pad4_16(v.Len())
}

func (v ChunkDATAValue) WriteTo(buf io.Writer) (int64, error) {
	binary.Write(buf, binary.BigEndian, v.ChunkDATAValueHeader)
	n := int64(12)
	n += int64(v.Size())
	buf.Write(v.Data)
	buf.Write(v.Padding)
	return n, nil
}

func (c ChunkDATA) IsUnordered() bool {
	return (c.Flags & 4) != 0
}

func (c *ChunkDATA) SetUnordered(v bool) {
	if v {
		c.Flags |= uint8(4)
	} else {
		c.Flags &= ^uint8(4)
	}
}

type CDATAFragmentState = uint8

const (
	FragMiddle CDATAFragmentState = iota
	FragEnd
	FragStart
	NoFrag
)

func (c ChunkDATA) FragmentState() CDATAFragmentState {
	return c.Flags & NoFrag
}

func (c *ChunkDATA) SetFragmentState(state CDATAFragmentState) {
	c.Flags >>= 2
	c.Flags <<= 2
	c.Flags |= (state & NoFrag)
}
