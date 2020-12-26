package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rickypai/sbench/fio/targets/crystaldiskmark"
)

var (
	directoryFlag = flag.String("directory", "", "directory of the test file")
	outputDirFlag = flag.String("outputdir", "", "output directory of results")
	loopsFlag     = flag.Int("loops", 1, "number of loops")
)

func main() {
	flag.Parse()

	if directoryFlag == nil || *directoryFlag == "" {
		log.Fatal("--directory flag required")
	}

	if outputDirFlag == nil || *outputDirFlag == "" {
		log.Fatal("--outputdir flag required")
	}

	ts := time.Now().Unix()

	outputBaseDir := filepath.Join(*outputDirFlag, strconv.FormatInt(ts, 10))
	err := os.MkdirAll(outputBaseDir, 0700)
	if err != nil {
		log.Fatal(err)
	}

	var stdout, stderr bytes.Buffer

	jobs := crystaldiskmark.CrystalDiskMarkTests(1, true)

	for _, job := range jobs {
		err = os.RemoveAll(*directoryFlag)
		if err != nil {
			log.Fatal(err)
		}

		err := os.MkdirAll(*directoryFlag, 0700)
		if err != nil {
			log.Fatal(err)
		}

		job.Loops = *loopsFlag

		err = job.Run(*directoryFlag, outputBaseDir, &stdout, &stderr)
		print(stdout.String())
		print(stderr.String())
		if err != nil {
			log.Fatal(err)
		}
	}
}
