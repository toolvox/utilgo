// 'codump' is a command-line utility that traverses a directory tree and aggregates the contents of files with specified extensions into a single output.

// Flags:
// -root: Specifies the root directory for the traversal. Defaults to the current directory if not specified.
// -exts: Defines a comma-separated list of file extensions. Only files with these extensions are included in the output.
// -out: Determines the output destination. If unspecified, Codump writes to standard output.

// The utility performs a recursive walk starting from the root directory, filtering files by the provided extensions. For each matching file, Codump appends the file's path and content to the output. This output can be directed to a file or standard output based on the user's choice.

// Example usage:
// To aggregate content from .go and .txt files into 'output.txt':
// codump -root . -exts ".go,.txt" -out output.txt
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"utilgo/pkg/flags"
)

// CodumpConfig holds configuration flags for processing the files.
type CodumpConfig struct {
	RootDir    string
	Extensions flags.CSVValue
	OutputFile flags.OutputFileValue
}

// processFiles processes files in the root directory matching the extensions and writes their contents to the output.
func (config *CodumpConfig) processFiles() error {
	writer, err := config.OutputFile.Writer()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error obtaining output writer: %v\n", err)
		os.Exit(1)
	}
	defer writer.Close()

	return filepath.Walk(config.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && slices.Contains(config.Extensions.Values, filepath.Ext(info.Name())) {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			if _, err := fmt.Fprintf(writer, "---\n%s\n\n", path); err != nil {
				return err
			}
			for scanner.Scan() {
				if _, err := fmt.Fprintln(writer, scanner.Text()); err != nil {
					return err
				}
			}
			return scanner.Err()
		}
		return nil
	})
}
