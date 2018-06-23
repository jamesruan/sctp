package format

type CommonHeader struct {
	SrcPort  uint16
	DstPort  uint16
	VTag     uint32
	CheckSum uint32 // CRC32c
}

type ChunkFieldHeader struct {
	Type   ChunkType
	Flags  uint8
	Length uint16
}

type ChunkField struct {
	ChunkFieldHeader
	Value []byte
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
