package main

import (
	"flag"
    "go/ast"
    "go/parser"
    "go/token"
	"go/printer"
	"log"
	"os"
	"errors"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	stringMethodName = "String"
)


func main(){
	var inputFile string
	var genFile string
	var outputFile string

	flag.StringVar(&genFile, "gen", "./gen/stringfunc.go", "path to the gen file")
	flag.StringVar(&inputFile, "input", "", "path to the input file")
	flag.StringVar(&outputFile, "output", "", "path to the output file. If non specifid, will override the input file.")
	flag.Parse()

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	if len(genFile) == 0 {
		log.Fatal("gen file is mandatory")
	}

	if len(outputFile) == 0 {
		outputFile = inputFile
	}

	genStringFunc, requiredImports, err := getStringFuncASTNode(genFile)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
	if err != nil {
		return
	}

	for _, reqiredImportPath := range requiredImports {
		astutil.AddImport(fset, f, reqiredImportPath)
	}

	astutil.Apply(f, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() != stringMethodName {
			return true
		}

		if len(funcDecal.Recv.List) != 1 {
			log.Fatal("invalid number of recievers")
		}

		genStringFunc.Recv = funcDecal.Recv
		// TODO: Allow for comments to be handled correctly
		genStringFunc.Doc = nil

		cr.Replace(genStringFunc)
		return false
	}, nil )

	if err := writeASTToFile(inputFile, fset, f); err != nil {
		log.Fatal(err)
	}
}


func getStringFuncASTNode(genfile string) (*ast.FuncDecl, []string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, genfile, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	var stringFuncNode *ast.FuncDecl


	requiredImports := getRequiredImports(fset, f)
	
	astutil.Apply(f, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() == stringMethodName {
			stringFuncNode = funcDecal
			return false
		}

		return true
	}, nil )

	if stringFuncNode == nil {
		return nil, nil, errors.New("Failed to find String Func")
	}

	return stringFuncNode, requiredImports, nil
}

func getRequiredImports(fset *token.FileSet, f *ast.File) []string {
	requiredImports := []string{}

	importLists := astutil.Imports(fset, f)
	for _, importList := range importLists {
		for _, importObj := range importList {
			path, err := strconv.Unquote(importObj.Path.Value)
			if err != nil {
				log.Fatal(err)
			}
			requiredImports = append(requiredImports, path)
		}
	}

	return requiredImports
}