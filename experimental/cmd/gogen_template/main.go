package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"log"
	"os"
	"strconv"

	"utilgo/pkg/flags"
)

//go:generate gen_test
type opts struct {
	File flags.FileValue
	Line uint

	Debug bool
}

func main() {
	var o opts
	flag.Var(&o.File, "file", "path to target file.")
	flag.UintVar(&o.Line, "line", 0, "line to start.")
	flag.BoolVar(&o.Debug, "debug", false, "use hard coded debug values.")
	flag.Parse()

	if o.Debug {
		o.File.Set("main.go")
		o.Line = 17
	}

	if len(o.File.String()) == 0 {
		genFile, okFile := os.LookupEnv("GOFILE")
		if !okFile {
			log.Printf("must provide file (by env or args)")
			os.Exit(1)
		}
		if err := o.File.Set(genFile); err != nil {
			log.Printf("set file option: %v", err)
			os.Exit(2)
		}
	}

	if o.Line == 0 {
		genLine, okLine := os.LookupEnv("GOLINE")
		if !okLine {
			log.Printf("must provide line (by env or args)")
			os.Exit(3)
		}

		line, err := strconv.Atoi(genLine)
		if err != nil {
			log.Printf("line '%s' in env to int: %v", genLine, err)
			os.Exit(4)
		}
		o.Line = uint(line) + 1
	}
	tokenFileSet := token.NewFileSet()
	fileContent := o.File.Get().([]byte)
	fileExpr, err := parser.ParseFile(tokenFileSet, o.File.Filename, fileContent, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		errs := err.(scanner.ErrorList)
		for _, e := range errs {
			log.Println(e)
		}
	}
	pos := tokenFileSet.File(1).LineStart(int(o.Line))
	endPos := pos
	for _, d := range fileExpr.Decls {
		if d.Pos() != pos {
			continue
		}
		endPos = d.End()
		break
	}

	if pos == endPos {
		log.Printf("no declaration at '%s:%d'", o.File.Filename, o.Line)
	}
	fmt.Println(string(fileContent[pos-1 : endPos]))
}
