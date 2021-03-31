package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"

	_ "embed"

	"golang.org/x/tools/go/ast/astutil"
)

const (
	stringFuncMethodName  = "String"
	stringFuncGenFileName = "./gen/stringfunc.go"
)

//go:embed gen/stringfunc.go
var stringFuncGenFile string

func getGenParseInfo() (*ParseInfo, error) {
	if len(stringFuncGenFile) == 0 {
		return nil, errors.New("Failed to parse string func file")
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, stringFuncGenFileName, stringFuncGenFile, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	return &ParseInfo{
		Fset: fset,
		F:    f,
	}, nil
}

func GenerateStringFunc(target *ParseInfo) error {
	genParseInfo, err := getGenParseInfo()
	if err != nil {
		return err
	}

	// Add missing imports to the target file for the newly generate func
	importsToAdd, err := getMissingImports(target, genParseInfo)
	if err != nil {
		return err
	}

	mutated := false

	astutil.Apply(target.F, func(cr *astutil.Cursor) bool {
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

		mutated = true

		for _, importToAdd := range importsToAdd {
			astutil.AddImport(target.Fset, target.F, importToAdd)
		}

		genStringFunc, err := getStringFuncASTNode(*funcDecal.Recv)
		if err != nil {
			log.Fatal(err)
		}

		cr.Replace(&genStringFunc)
		return false
	}, nil)

	if mutated {
		ast.SortImports(target.Fset, target.F)
	}

	return nil
}

func getMissingImports(target, genParseInfo *ParseInfo) ([]string, error) {
	genRequiredImports, err := getImports(genParseInfo)
	if err != nil {
		return nil, err
	}

	importsToAdd := []string{}
	for genImport := range genRequiredImports {
		if !astutil.UsesImport(target.F, genImport) {
			importsToAdd = append(importsToAdd, genImport)
		}
	}

	return importsToAdd, nil
}

func getNonPointTypeFromRecv(recv ast.FieldList) ast.Expr {
	starExpr, ok := recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return recv.List[0].Type
	}

	return starExpr.X
}

func getStringFuncASTNode(targetRecv ast.FieldList) (ast.FuncDecl, error) {
	var stringFuncNode *ast.FuncDecl

	genParseInfo, err := getGenParseInfo()
	if err != nil {
		return ast.FuncDecl{}, nil
	}

	astutil.Apply(genParseInfo.F, func(cr *astutil.Cursor) bool {
		switch node := cr.Node().(type) {
		case *ast.FuncDecl:
			if node.Name.String() == stringFuncMethodName {
				stringFuncNode = node
			}
		// Handle replacing type for copy declaration
		// TODO: should only be scoped to String func instead of global
		case *ast.DeclStmt:
			genDecl, ok := node.Decl.(*ast.GenDecl)
			if !ok {
				return true
			}

			if len(genDecl.Specs) != 1 {
				return true
			}

			valSpec, ok := genDecl.Specs[0].(*ast.ValueSpec)
			if !ok {
				return true
			}

			if valSpec.Names[0].String() != "copy" {
				return true
			}

			valSpec.Type = getNonPointTypeFromRecv(targetRecv)
			return false
		}
		return true
	}, nil)

	if stringFuncNode == nil {
		return ast.FuncDecl{}, errors.New("Failed to find String Func")
	}

	retStringFunc := addRecv(*stringFuncNode, targetRecv)

	return retStringFunc, nil
}

// addRecv adds a receiver to a func. Everything passed by value to ensure no side effects
func addRecv(funcDecl ast.FuncDecl, recv ast.FieldList) ast.FuncDecl {
	funcDecl.Recv = &recv
	return funcDecl
}

func getImports(target *ParseInfo) (map[string]bool, error) {
	requiredImports := map[string]bool{}

	importLists := astutil.Imports(target.Fset, target.F)
	for _, importList := range importLists {
		for _, importObj := range importList {
			path, err := strconv.Unquote(importObj.Path.Value)
			if err != nil {
				return nil, err
			}
			requiredImports[path] = true
		}
	}

	return requiredImports, nil
}
