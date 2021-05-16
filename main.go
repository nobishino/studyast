package main

import (
	_ "embed"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

//go:embed main.go
var src string

type MyInt int
type MyMyInt MyInt
type AliasInt = int
type IntSlice []int

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
	}
	pkg, err := cfg.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		log.Fatal(err)
	}
	underlyingIsItself := func(typeName string) {
		tp := pkg.Scope().Lookup(typeName)
		if tp == nil {
			log.Printf("type %q not found\n", typeName)
			return
		}
		result := tp.Type() == tp.Type().Underlying()
		fmt.Printf("%[1]sのunderlying typeは%[1]s? -> %t\n", typeName, result)
	}
	underlyingIsItself("MyInt")
	underlyingIsItself("MyMyInt")
	underlyingIsItself("AliasInt")
	underlyingIsItself("IntSlice")
}
