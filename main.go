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


// Will not redact non exported, but using json to convert will remove unexported fields
// Handle Comments


func main(){
	var inputFile string
	var genFile string
	flag.StringVar(&genFile, "gen", "./gen/stringfunc.go", "path to the gen file")
	flag.StringVar(&inputFile, "input", "", "path to input file")
	flag.Parse()

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	if len(genFile) == 0 {
		log.Fatal("gen file is mandatory")
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
	if err != nil {
		return
	}

	genStringFunc, requiredImports, err := getStringFuncASTNode(genFile)
	if err != nil {
		log.Fatal(err)
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

func writeASTToFile(filename string,  fset *token.FileSet, f *ast.File) error {
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