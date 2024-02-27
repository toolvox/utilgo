package main

import (
	"flag"
	"fmt"
	"os"

	"utilgo/pkg/flags"
)

func main() {
	config := CodumpConfig{
		OutputFile: flags.NewOutputFileValue("", os.O_CREATE|os.O_TRUNC),
	}
	flag.StringVar(&config.RootDir, "root", ".", "Root directory to walk through")
	flag.Var(&config.Extensions, "exts", "Comma-separated list of file extensions to include (e.g., .go,.txt)")
	flag.Var(&config.OutputFile, "out", "Output file path. Defaults to stdout if empty")
	flag.Parse()

	if err := config.processFiles(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
