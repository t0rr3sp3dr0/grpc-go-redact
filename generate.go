package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	stringFuncMethodName = "String"
	stringFuncGenFile    = "./gen/stringfunc.go"
)

func GenerateStringFunc(fset *token.FileSet, f *ast.File) error {
	genStringFunc, requiredImports, err := getStringFuncASTNode(stringFuncGenFile)
	if err != nil {
		return err
	}

	// Add required imports from the pre-gen file
	for _, reqiredImportPath := range requiredImports {
		astutil.AddImport(fset, f, reqiredImportPath)
	}

	astutil.Apply(f, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() != stringFuncMethodName {
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
	}, nil)

	return nil
}

func getStringFuncASTNode(genfile string) (*ast.FuncDecl, []string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, genfile, nil, parser.ParseComments)
	if err != nil {
		return nil, nil, err
	}

	var stringFuncNode *ast.FuncDecl

	requiredImports, err := getRequiredImports(fset, f)
	if err != nil {
		return nil, nil, err
	}

	astutil.Apply(f, func(cr *astutil.Cursor) bool {
		funcDecal, ok := cr.Node().(*ast.FuncDecl)
		if !ok {
			return true
		}
		if funcDecal.Name.String() == stringFuncMethodName {
			stringFuncNode = funcDecal
			return false
		}

		return true
	}, nil)

	if stringFuncNode == nil {
		return nil, nil, errors.New("Failed to find String Func")
	}

	return stringFuncNode, requiredImports, nil
}

func getRequiredImports(fset *token.FileSet, f *ast.File) ([]string, error) {
	requiredImports := []string{}

	importLists := astutil.Imports(fset, f)
	for _, importList := range importLists {
		for _, importObj := range importList {
			path, err := strconv.Unquote(importObj.Path.Value)
			if err != nil {
				return nil, err
			}
			requiredImports = append(requiredImports, path)
		}
	}

	return requiredImports, nil
}
