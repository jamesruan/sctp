package utils

type Serial32 struct {
	uint32
}

const two31 = 1 << 31

// Add n to s. n should be in range 0..2^31-1, or a panic will occur.
func (s Serial32) Add(n int) Serial32 {
	if n > two31-1 {
		panic("add a number too large")
	}
	return Serial32{uint32((int(s.uint32) + n) & (1<<32 - 1))}
}

func (s Serial32) ToUInt32() uint32 {
	return s.uint32
}

func MakeSerial32(s uint32) Serial32 {
	return Serial32{s}
}

func (s Serial32) Eq(t Serial32) bool {
	return s.uint32 == t.uint32
}

func (s Serial32) LessThan(t Serial32) bool {
	if s.Eq(t) {
		return false
	}
	if s.uint32 < t.uint32 {
		return t.uint32-s.uint32 < two31
	} else {
		return s.uint32-t.uint32 > two31
	}
}

func (s Serial32) GreaterThan(t Serial32) bool {
	if s.Eq(t) {
		return false
	}
	if s.uint32 < t.uint32 {
		return t.uint32-s.uint32 > two31
	} else {
		return s.uint32-t.uint32 < two31
	}
}
