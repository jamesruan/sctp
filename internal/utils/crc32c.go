package utils

import (
	"hash"
	"hash/crc32"
)

// Hash32.Sum() returns in big-endian
func NewCRC32c() hash.Hash32 {
	crc32c_table := crc32.MakeTable(crc32.Castagnoli)
	return crc32.New(crc32c_table)
}
