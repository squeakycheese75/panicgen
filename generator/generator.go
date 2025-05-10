package generator

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed test.tmpl
var testTemplate string

type FuncData struct {
	FuncName string
	Package  string
}

func GenerateTests(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for pkgName, pkg := range pkgs {
		var funcs []FuncData

		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if ok && fn.Recv == nil && fn.Name.IsExported() {
					funcs = append(funcs, FuncData{
						FuncName: fn.Name.Name,
						Package:  pkgName,
					})
				}
			}
		}

		if len(funcs) > 0 {
			return writeTestFile(dir, funcs)
		}
	}

	return nil
}

func writeTestFile(dir string, funcs []FuncData) error {
	// Load the embedded template
	tmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(dir, "panic_test.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, funcs)
}
