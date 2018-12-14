package main

import (
	"fmt"
	"os"
	"text/template"
)

// TypeContext represents information needed to perform code generatoin
// for default types.
type TypeContext struct {
	Name   string
	GoType string
}

func codeGenerate(ctx TypeContext, filepath string) error {
	f, err := os.Create(filepath)
	defer f.Close()

	if err != nil {
		return err
	}

	t, err := template.New("collection.tmpl").ParseFiles("codegen/template/collection.tmpl")
	if err != nil {
		return err
	}

	if err = t.Execute(f, ctx); err != nil {
		return err
	}

	return nil
}

func main() {
	ctxs := []TypeContext{
		TypeContext{Name: "Int", GoType: "int"},
		TypeContext{Name: "String", GoType: "string"},
	}

	for _, ctx := range ctxs {
		path := fmt.Sprintf("pkg/collections/%s_list.go", ctx.GoType)

		fmt.Printf("Generating code for type %s at %s\n", ctx.Name, path)

		codeGenerate(ctx, path)
	}
}
