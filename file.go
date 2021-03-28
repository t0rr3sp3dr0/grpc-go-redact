package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	generatedFileExtension = ".pb.go"
)

type ParseInfo struct {
	OutputFile string
	Fset       *token.FileSet
	F          *ast.File
}

func protoGenFileFilter(fileInfo fs.FileInfo) bool {
	fmt.Println(fileInfo.Name())
	return strings.HasSuffix(fileInfo.Name(), generatedFileExtension)
}

func ParseFile(filePath string) (*ParseInfo, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s with err: %v", filePath, err)
	}

	return &ParseInfo{
		Fset: fset,
		F:    f,
	}, nil
}

func ParseDir(dirPath string) ([]*ParseInfo, error) {
	parseInfos := []*ParseInfo{}

	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !strings.HasSuffix(path, generatedFileExtension) {
				return nil
			}

			fmt.Println("Parsing file: ", path)

			parseInfo, err := ParseFile(path)
			if err != nil {
				return err
			}

			parseInfos = append(parseInfos, parseInfo)

			return nil
		})
	if err != nil {
		return nil, err
	}

	return parseInfos, nil
}

func writeASTToFile(parseInfo *ParseInfo) error {
	file, err := os.OpenFile(parseInfo.OutputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := printer.Fprint(file, parseInfo.Fset, parseInfo.F); err != nil {
		return err
	}

	return nil
}
