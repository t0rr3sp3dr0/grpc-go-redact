package main

import (
	"context"
	"flag"
	"github.com/Azure/grpc-go-redact/filehandler"
	"github.com/Azure/grpc-go-redact/worker"

	"github.com/samkreter/go-core/log"
)

const (
	defaultNumWorkers = 5
)

func main() {
	var inputFile string
	var outputFile string
	var inputDir string
	var numWorkers int
	var verboseLogging bool

	flag.BoolVar(&verboseLogging, "v", false, "enable verbose logging")
	flag.IntVar(&numWorkers, "workers", defaultNumWorkers, "the number of concurrent workers")
	flag.StringVar(&inputFile, "input", "", "path to the input file")
	flag.StringVar(&inputDir, "dir", "", "path to the input dir")
	flag.StringVar(&outputFile, "output", "", "path to the output file. If non specifid, will override the input file.")
	flag.Parse()

	ctx := context.Background()
	logger := log.G(ctx)

	if len(inputFile) == 0 && len(inputDir) == 0 {
		logger.Fatal("input file or dir is mandatory")
	}

	logLevel := "info"
	if verboseLogging {
		logLevel = "debug"
	}

	if err := log.SetLogLevel(logLevel); err != nil {
		logger.Errorf("failed to set log level to : '%s'", logLevel)
	}

	if len(outputFile) == 0 {
		outputFile = inputFile
	}

	fileToGenerate := []*filehandler.ParseInfo{}

	// Handle single file
	if len(inputFile) != 0 {
		parseInfo, err := filehandler.ParseFile(inputFile)
		if err != nil {
			logger.Fatal(err)
		}

		parseInfo.OutputFile = outputFile
		fileToGenerate = append(fileToGenerate, parseInfo)
	}

	// Handle parsing entire dir
	if len(inputDir) != 0 {
		parseInfos, err := filehandler.ParseDir(inputDir)
		if err != nil {
			logger.Fatal(err)
		}

		fileToGenerate = append(fileToGenerate, parseInfos...)
	}

	workQueue := worker.NewWorkQueue(len(fileToGenerate))
	defer workQueue.Shutdown()

	for _, target := range fileToGenerate {
		workQueue.AddJob(target)
	}

	if len(fileToGenerate) < numWorkers {
		numWorkers = len(fileToGenerate)
	}

	logger.Debugln("Number of jobs to process: ", workQueue.NumJobs())
	logger.Debugln("Starting ", numWorkers, " workers")

	workQueue.StartWorkers(numWorkers)
	workQueue.WaitForJobs()
}
