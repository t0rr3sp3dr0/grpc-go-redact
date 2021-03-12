package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type ParseInfo struct {
	Fset *token.FileSet
	F    *ast.File
}

func writeASTToFile(filename string, parseInfo *ParseInfo) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := printer.Fprint(file, parseInfo.Fset, parseInfo.F); err != nil {
		return err
	}

	return nil
}
