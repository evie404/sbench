package crystaldiskmark

import (
	"fmt"

	"github.com/rickypai/sbench/fio"
	"github.com/rickypai/sbench/fio/defaults"
)

const (
	crystalDiskMarkDefaultRWMixRead = 70
)

func CrystalDiskMarkTests(loops int, includeMixed bool) []fio.Job {
	jobs := make([]fio.Job, 0, 12)

	for _, rw := range fio.AllRWs(includeMixed) {
		jobs = append(
			jobs,
			[]fio.Job{
				baseJob(rw, fio.AccessSequential, "1m", 8, 1, loops),
				baseJob(rw, fio.AccessSequential, "1m", 1, 1, loops),
				baseJob(rw, fio.AccessRandom, "4k", 32, 1, loops),
				baseJob(rw, fio.AccessRandom, "4k", 1, 1, loops),
			}...,
		)
	}

	return jobs
}

func CrystalDiskMarkNVMeTests(loops int, includeMixed bool) []fio.Job {
	jobs := make([]fio.Job, 0, 12)

	for _, rw := range fio.AllRWs(includeMixed) {
		jobs = append(
			jobs,
			[]fio.Job{
				baseJob(rw, fio.AccessSequential, "1m", 8, 1, loops),
				baseJob(rw, fio.AccessSequential, "128k", 32, 1, loops),
				baseJob(rw, fio.AccessRandom, "4k", 32, 16, loops),
				baseJob(rw, fio.AccessRandom, "4k", 1, 1, loops),
			}...,
		)
	}

	return jobs
}

func baseJob(rw fio.RW, mode fio.Access, blockSize string, depth, jobs, loops int) fio.Job {
	name := fmt.Sprintf("%s-%s-%s-q%vt%v", rw, mode, blockSize, depth, jobs)

	return fio.Job{
		Name:      name,
		BlockSize: blockSize,
		IODepth:   depth,
		NumJobs:   jobs,
		ReadWrite: fio.NewRWAccess(rw, mode),
		Loops:     loops,

		RWMixRead: crystalDiskMarkDefaultRWMixRead,

		Direct:   defaults.DefaultDirect,
		Size:     defaults.DefaultSize,
		IOEngine: defaults.DefaultIOEngine,
	}
}
