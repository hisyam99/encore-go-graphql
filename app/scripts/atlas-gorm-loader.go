package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"encore.app/app"
)

// Define the models to generate migrations for.
var models = []any{
	&app.User{},
	&app.Category{},
	&app.ResumeContent{},
	&app.Project{},
	&app.Blog{},
}

func main() {
	stmts, err := gormschema.New("postgres").Load(models...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
