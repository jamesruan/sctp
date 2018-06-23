package format

import (
	"bytes"
	"encoding/binary"
)

type ChunkFieldHeader struct {
	Type   ChunkType
	Flags  uint8
	Length uint16
}

type ChunkField struct {
	ChunkFieldHeader
	Params []ChunkParam
	Value  []byte
}

func (s ChunkField) MarshalBinary() (data []byte, err error) {
	paddedLen := 4 - (len(s.Value) & 3)
	paddingBytes := make([]byte, paddedLen)
	buf := bytes.NewBuffer(make([]byte, 0, int(s.Length)+paddedLen))

	binary.Write(buf, binary.BigEndian, s.ChunkFieldHeader)
	buf.Write(s.Value)
	buf.Write(paddingBytes)
	return buf.Bytes(), err
}

func (s *ChunkFieldHeader) UnmarshalBinary(data []byte) (err error) {
	r := bytes.NewReader(data)
	return binary.Read(r, binary.BigEndian, s)
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
	Value []byte
}

func (s ChunkParam) MarshalBinary() (data []byte, err error) {
	paddedLen := 4 - (len(s.Value) & 3)
	paddingBytes := make([]byte, paddedLen)
	buf := bytes.NewBuffer(make([]byte, 0, int(s.Length)+paddedLen))

	binary.Write(buf, binary.BigEndian, s.ChunkParamHeader)
	buf.Write(s.Value)
	buf.Write(paddingBytes)
	return buf.Bytes(), err
}

func (s *ChunkParamHeader) UnmarshalBinary(data []byte) (err error) {
	r := bytes.NewReader(data)
	return binary.Read(r, binary.BigEndian, s)
}

type ChunkParamType = uint16

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
