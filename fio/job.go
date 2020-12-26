package fio

type Job struct {
	Name      string
	Direct    bool
	IODepth   int
	NumJobs   int
	RWMixRead int
	IOEngine  IOEngine
	BlockSize string
	Size      string
	Directory string
	Loops     int
	ReadWrite RWAccess
	Thread    bool
}
