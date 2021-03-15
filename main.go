package main

import (
	"flag"
    "go/parser"
    "go/token"
	"log"
)



func main(){
	var inputFile string
	var outputFile string

	flag.StringVar(&inputFile, "input", "", "path to the input file")
	flag.StringVar(&outputFile, "output", "", "path to the output file. If non specifid, will override the input file.")
	flag.Parse()

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}


	if len(outputFile) == 0 {
		outputFile = inputFile
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputFile, nil, parser.AllErrors)
	if err != nil {
		return
	}

	target := &ParseInfo{
		Fset: fset,
		F: f,
	}

	if err := GenerateStringFunc(target); err != nil {
		log.Fatal(err)
	}

	if err := writeASTToFile(outputFile, target); err != nil {
		log.Fatal(err)
	}
}




