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

func main() {
	// main.go 自身を読み込んでAST(*ast.File)を作る
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}
	// ASTに対して型チェックを行う
	cfg := &types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
	}
	pkg, err := cfg.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		log.Fatal(err)
	}
	// 与えられた型名に対して、「underlying typeが自分自身かどうか」をチェックする関数
	underlyingIsItself := func(typeName string) {
		tp := pkg.Scope().Lookup(typeName)
		if tp == nil {
			log.Printf("type %q not found\n", typeName)
			return
		}
		result := types.Identical(tp.Type(), tp.Type().Underlying())
		fmt.Printf("%[1]sのunderlying typeは%[1]s? -> %t\n", typeName, result)
	}
	underlyingIsItself("MyInt")    // MyIntのunderlying typeはMyInt? -> false
	underlyingIsItself("MyMyInt")  // MyMyIntのunderlying typeはMyMyInt? -> false
	underlyingIsItself("AliasInt") // AliasIntのunderlying typeはAliasInt? -> true
	underlyingIsItself("IntSlice") // IntSliceのunderlying typeはIntSlice? -> false
}

type MyInt int
type MyMyInt MyInt
type AliasInt = int
type IntSlice []int
