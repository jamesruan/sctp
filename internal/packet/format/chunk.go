package format

import (
	"encoding/binary"
	"io"
)

var padding4 = [3]byte{0}

type ChunkFieldHeader struct {
	Type   ChunkType
	Flags  uint8
	Length uint16
}

type ChunkField struct {
	ChunkFieldHeader
	Value io.WriterTo
}

func (c ChunkField) Len() uint16 {
	return c.ChunkFieldHeader.Length
}

func (c ChunkField) Size() uint16 {
	return pad4_16(c.Len())
}

func (s ChunkField) WriteTo(buf io.Writer) (int64, error) {
	binary.Write(buf, binary.BigEndian, s.ChunkFieldHeader)
	n, err := s.Value.WriteTo(buf)
	return n + 4, err
}

type ChunkType = uint8

const (
	CT_DATA ChunkType = iota
	CT_INIT
	CT_INIT_ACK
	CT_SACK
	CT_HEARTBEAT
	CT_HEARTBEAT_ACK
	CT_ABORT
	CT_SHUTDOWN
	CT_SHUTDOWN_ACK
	CT_ERROR
	CT_COOKIE_ECHO
	CT_COOKIE_ACK
	CT_ECNE
	CT_CWR
	CT_SHUTDOWN_COMPLETE
)

type UnknownChunkFieldAction uint8

const (
	UCFA_DiscardPacket UnknownChunkFieldAction = iota
	UCFA_DiscardPacketAndReport
	UCFA_SkipField
	UCFA_SkipFieldAndReport
)

func (c ChunkFieldHeader) GetUnknownChunkFieldAction() UnknownChunkFieldAction {
	switch (c.Type >> 6) & 3 {
	case 0:
		return UCFA_DiscardPacket
	case 1:
		return UCFA_DiscardPacketAndReport
	case 2:
		return UCFA_SkipField
	default:
		return UCFA_SkipFieldAndReport
	}
}

type ChunkParamHeader struct {
	Type   ChunkParamType
	Length uint16
}

type ChunkParam struct {
	ChunkParamHeader
	Value io.WriterTo
}

func (s ChunkParam) WriteTo(buf io.Writer) (int64, error) {
	binary.Write(buf, binary.BigEndian, s.ChunkParamHeader)
	n, err := s.Value.WriteTo(buf)
	return n + 4, err
}

type ChunkParamType = uint16

const (
	CPT_IPv4Addr           ChunkParamType = 5
	CPT_IPv6Addr                          = 6
	CPT_CookiePreservative                = 9
	CPT_ECNCapable                        = 0x8000
	CPT_HostNameAddr                      = 11
	CPT_SupportedAddrTypes                = 12
)

type ChunkParamAddrType = uint16

const (
	CPAT_IPv4Addr     ChunkParamAddrType = 5
	CPAT_IPv6Addr                        = 6
	CPAT_HostNameAddr                    = 11
)

type UnknownChunkParamAction uint8

const (
	UCPA_DiscardChunk UnknownChunkParamAction = iota
	UCPA_DiscardChunkAndReport
	UCPA_SkipParam
	UCPA_SkipParamAndReport
)

func (c ChunkParamHeader) GetUnknownChunkParamAction() UnknownChunkParamAction {
	switch (c.Type >> 14) & 3 {
	case 0:
		return UCPA_DiscardChunk
	case 1:
		return UCPA_DiscardChunkAndReport
	case 2:
		return UCPA_SkipParam
	default:
		return UCPA_SkipParamAndReport
	}
}

func pad4_16(v uint16) uint16 {
	return v + (4 - (v & 3))
}
