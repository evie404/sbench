package fio

type IOEngine string

const (
	IOEnginePosixAIO   IOEngine = "posixaio"
	IOEngineLibAIO     IOEngine = "libaio"
	IOEngineWindowsAIO IOEngine = "windowsaio"
)

type Access string

const (
	AccessSequential Access = "seq"
	AccessRandom     Access = "rand"
)

func AllAccesses() []Access {
	return []Access{
		AccessSequential,
		AccessRandom,
	}
}

type RW string

const (
	RWRead  RW = "read"
	RWWrite RW = "write"
	RWMix   RW = "mix"
)

func AllRWs(includeMix bool) []RW {
	rws := []RW{
		RWRead,
		RWWrite,
	}

	if includeMix {
		rws = append(rws, RWMix)
	}

	return rws
}

type RWAccess string

const (
	RWAccessSequentialRead  RWAccess = "read"
	RWAccessSequentialWrite RWAccess = "write"
	RWAccessSequentialMixed RWAccess = "readwrite"

	RWAccessRandomRead  RWAccess = "randread"
	RWAccessRandomWrite RWAccess = "randwrite"
	RWAccessRandomMixed RWAccess = "randrw"

	RWAccessUnknown RWAccess = "unknown"
)

func NewRWAccess(rw RW, mode Access) RWAccess {
	if mode == AccessSequential {
		switch rw {
		case RWRead:
			return RWAccessSequentialRead
		case RWWrite:
			return RWAccessSequentialWrite
		case RWMix:
			return RWAccessSequentialMixed
		}
	}

	switch rw {
	case RWRead:
		return RWAccessRandomRead
	case RWWrite:
		return RWAccessRandomWrite
	case RWMix:
		return RWAccessRandomMixed
	}

	return RWAccessUnknown
}
