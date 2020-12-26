package fio

import (
	"io"
	"os/exec"
	"strconv"
)

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

func (j *Job) Run(stdout, stderr io.Writer) error {
	direct := "0"
	if j.Direct {
		direct = "1"
	}

	thread := "0"
	if j.Thread {
		thread = "1"
	}

	cmd := exec.Command(
		"fio",
		"--name="+j.Name,
		"--direct="+direct,
		"--iodepth="+strconv.Itoa(j.IODepth),
		"--numjobs="+strconv.Itoa(j.NumJobs),
		"--max-jobs="+strconv.Itoa(j.NumJobs),
		"--rwmixread="+strconv.Itoa(j.RWMixRead),
		"--ioengine="+string(j.IOEngine),
		"--blocksize="+j.BlockSize,
		"--size="+j.Size,
		"--loops="+strconv.Itoa(j.Loops),
		"--rw="+string(j.ReadWrite),
		"--thread="+thread,

		"--parse-only",

		"--directory="+j.Directory,
	)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
