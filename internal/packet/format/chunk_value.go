package format

import (
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrInvalidValue = errors.New("invalid value")
)

// Chunk DATA

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

// Chunk INIT

type ChunkINIT struct {
	ChunkFieldHeader
	ChunkINITValue
}

type ChunkINITValue struct {
	I_Tag  uint32 // Initial Tag
	A_RWND uint32 // advertised receiver window
	OS     uint16 // number of Outbound Streams
	MIS    uint16 // number of Inbound Streams
	I_TSN  uint32
	Params []ChunkParam
}

// CPT_IPv4Addr
type ChunkParamIPv4AddrValue struct {
	Addr [4]byte
}

func (c ChunkParamIPv4AddrValue) WriteTo(buf io.Writer) (int64, error) {
	n, err := buf.Write(c.Addr[:])
	return int64(n), err
}

// CPT_IPv6Addr
type ChunkParamIPv6AddrValue struct {
	Addr [16]byte
}

func (c ChunkParamIPv6AddrValue) WriteTo(buf io.Writer) (int64, error) {
	n, err := buf.Write(c.Addr[:])
	return int64(n), err
}

// CPT_CookiePreservative
type ChunkParamCookiePreservative struct {
	LSI uint32 //suggested Life-Span Increment (msec.)
}

func (c ChunkParamCookiePreservative) WriteTo(buf io.Writer) (int64, error) {
	if err := binary.Write(buf, binary.BigEndian, c.LSI); err != nil {
		return 0, err
	} else {
		return 4, nil
	}
}

// CPT_HostNameAddr
type ChunkParamHostNameAddr struct {
	HostName []byte
}

func (c ChunkParamHostNameAddr) WriteTo(buf io.Writer) (int64, error) {
	n, err := buf.Write(c.HostName[:])
	return int64(n), err
}

// CPT_SupportedAddrTypes
type ChunkParamSupportedAddrTypes struct {
	AddrTypes []ChunkParamAddrType
}

func (c ChunkParamSupportedAddrTypes) WriteTo(buf io.Writer) (int64, error) {
	if err := binary.Write(buf, binary.BigEndian, c.AddrTypes); err != nil {
		return 0, err
	} else {
		return int64(2 * len(c.AddrTypes)), nil
	}
}
