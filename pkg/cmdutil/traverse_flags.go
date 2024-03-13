// Package cmdutil provides helpers and utilities for cmd apps.
package cmdutil

import (
	"flag"
	"io/fs"
	"os"
	"slices"
	"strings"

	"utilgo/pkg/flags"
	"utilgo/pkg/reflectutil"
)

// TraverseFlags encapsulates options for file system traversal and pattern matching.
type TraverseFlags struct {
	RootDir      string
	IncludeGlobs flags.CSVValue
	ExcludeGlobs flags.CSVValue
	IgnoreFile   string
}

// RegisterVars defines command-line flags for configuring file system traversal.
func (o *TraverseFlags) RegisterVars(targetSet *flag.FlagSet) {
	targetSet.StringVar(&o.RootDir, "root", ".", "Set the root directory for traversal.")
	targetSet.Var(&o.IncludeGlobs, "include", "Glob patterns for files to include. Defaults to all files.")
	targetSet.Var(&o.ExcludeGlobs, "exclude", "Glob patterns for files to exclude. Defaults to no files.")
	targetSet.StringVar(&o.IgnoreFile, "ignore", "", "Path to an ignore file (e.g., .gitignore) for additional exclude patterns.")
}

// RootFS creates an fs.FS representing the root directory for traversal.
func (o *TraverseFlags) RootFS() fs.FS {
	f := os.DirFS(o.RootDir)
	o.RootDir = "."
	return f
}

// Ignores reads an ignore file, if specified, parsing it into a slice of exclude patterns.
func (o TraverseFlags) Ignores(f fs.FS) ([]string, error) {
	if o.IgnoreFile == "" {
		return nil, nil
	}

	ignores, err := fs.ReadFile(f, o.IgnoreFile)
	if err != nil {
		return nil, err
	}

	var ignoreExcludes []string
	for _, line := range strings.Split(string(ignores), "\n") {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			ignoreExcludes = append(ignoreExcludes, trimmed)
		}
	}
	
	ignoreExcludes = append(ignoreExcludes, o.IgnoreFile)
	return slices.DeleteFunc(ignoreExcludes, reflectutil.IsZero), nil
}
