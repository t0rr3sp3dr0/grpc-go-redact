package main

import (
	"flag"
	"log"
)

const (
	defaultNumWorkers = 5
)

func main() {
	var inputFile string
	var outputFile string
	var inputDir string
	var numWorkers int

	flag.IntVar(&numWorkers, "workers", defaultNumWorkers, "the number of concurrent workers")
	flag.StringVar(&inputFile, "input", "", "path to the input file")
	flag.StringVar(&inputDir, "dir", "", "path to the input dir")
	flag.StringVar(&outputFile, "output", "", "path to the output file. If non specifid, will override the input file.")
	flag.Parse()

	if len(inputFile) == 0 && len(inputDir) == 0 {
		log.Fatal("input file or dir is mandatory")
	}

	if len(outputFile) == 0 {
		outputFile = inputFile
	}

	fileToGenerate := []*ParseInfo{}

	// Handle single file
	if len(inputFile) != 0 {
		parseInfo, err := ParseFile(inputFile)
		if err != nil {
			log.Fatal(err)
		}

		parseInfo.OutputFile = outputFile
		fileToGenerate = append(fileToGenerate, parseInfo)
	}

	// Handle parsing entire dir
	if len(inputDir) != 0 {
		parseInfos, err := ParseDir(inputDir)
		if err != nil {
			log.Fatal(err)
		}

		fileToGenerate = append(fileToGenerate, parseInfos...)
	}

	workQueue := NewWorkQueue(len(fileToGenerate))
	defer workQueue.Shutdown()

	for _, target := range fileToGenerate {
		workQueue.AddJob(target)
	}

	if len(fileToGenerate) < numWorkers {
		numWorkers = len(fileToGenerate)
	}

	log.Println("Number of jobs to process: ", workQueue.NumJobs())
	log.Println("Starting ", numWorkers, " workers")

	workQueue.StartWorkers(numWorkers)
	workQueue.WaitForJobs()
}
