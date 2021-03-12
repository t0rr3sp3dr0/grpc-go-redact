package main

import (
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func writeASTToFile(filename string, fset *token.FileSet, f *ast.File) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := printer.Fprint(file, fset, f); err != nil {
		return err
	}

	return nil
}
