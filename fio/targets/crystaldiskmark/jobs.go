package crystaldiskmark

import (
	"fmt"

	"github.com/rickypai/sbench/consts"
	"github.com/rickypai/sbench/fio"
	"github.com/rickypai/sbench/fio/defaults"
)

const (
	crystalDiskMarkDefaultRWMixRead = 70
)

func CrystalDiskMarkTests(storageType consts.StorageType, loops int, includeMixed bool) []fio.Job {
	jobs := make([]fio.Job, 0, 12)

	for _, rw := range fio.AllRWs(includeMixed) {
		jobs = append(
			jobs,
			[]fio.Job{
				baseJob(storageType, rw, fio.AccessSequential, "1m", 8, 1, loops),
				baseJob(storageType, rw, fio.AccessSequential, "1m", 1, 1, loops),
				baseJob(storageType, rw, fio.AccessRandom, "4k", 32, 1, loops),
				baseJob(storageType, rw, fio.AccessRandom, "4k", 1, 1, loops),
			}...,
		)
	}

	return jobs
}

func CrystalDiskMarkNVMeTests(storageType consts.StorageType, loops int, includeMixed bool) []fio.Job {
	jobs := make([]fio.Job, 0, 12)

	for _, rw := range fio.AllRWs(includeMixed) {
		jobs = append(
			jobs,
			[]fio.Job{
				baseJob(storageType, rw, fio.AccessSequential, "1m", 8, 1, loops),
				baseJob(storageType, rw, fio.AccessSequential, "128k", 32, 1, loops),
				baseJob(storageType, rw, fio.AccessRandom, "4k", 32, 16, loops),
				baseJob(storageType, rw, fio.AccessRandom, "4k", 1, 1, loops),
			}...,
		)
	}

	return jobs
}

func testSize(storageType consts.StorageType, rw fio.RW, mode fio.Access) string {
	switch storageType {
	case consts.StorageTypeNVMe:
		return nvmeTestSize(rw, mode)
	case consts.StorageTypeSSD:
		return ssdTestSize(rw, mode)
	case consts.StorageTypeUSB:
		return usbTestSize(rw, mode)
	case consts.StorageTypeHDD:
		return hddTestSize(rw, mode)
	}

	return hddTestSize(rw, mode)
}

func usbTestSize(rw fio.RW, mode fio.Access) string {
	if mode == fio.AccessRandom {
		if rw == fio.RWRead {
			return "128m"
		}

		return "64m"
	}

	// sequential
	if rw == fio.RWRead {
		return "1g"
	}

	return "512m"
}

func hddTestSize(rw fio.RW, mode fio.Access) string {
	return usbTestSize(rw, mode)
}

func ssdTestSize(rw fio.RW, mode fio.Access) string {
	if mode == fio.AccessRandom {
		return "1g"
	}

	return "2g"
}

func nvmeTestSize(rw fio.RW, mode fio.Access) string {
	if mode == fio.AccessRandom {
		return "2g"
	}

	return "4g"
}

func baseJob(storageType consts.StorageType, rw fio.RW, mode fio.Access, blockSize string, depth, jobs, loops int) fio.Job {
	name := fmt.Sprintf("%s-%s-%s-q%vt%v", rw, mode, blockSize, depth, jobs)
	size := testSize(storageType, rw, mode)

	return fio.Job{
		Name:      name,
		BlockSize: blockSize,
		IODepth:   depth,
		NumJobs:   jobs,
		ReadWrite: fio.NewRWAccess(rw, mode),
		Loops:     loops,

		RWMixRead: crystalDiskMarkDefaultRWMixRead,

		Direct:   defaults.DefaultDirect,
		Size:     size,
		IOEngine: defaults.DefaultIOEngine,
		Runtime:  defaults.DefaultRuntime,
	}
}
