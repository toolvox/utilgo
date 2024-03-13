// Package files provides helpers  and utilities for working with Files and FileSystems.
package files

import (
	"io/fs"
	"path/filepath"
	"slices"

	"utilgo/pkg/errs"
)

// GlobMatcher holds patterns for inclusion and exclusion of file paths.
// Directories don't need to be included but can be excluded.
type GlobMatcher struct {
	Include []string
	Exclude []string
}

// NewGlobMatcher constructs a new [GlobMatcher] with specified include and exclude patterns.
func NewGlobMatcher(includes, excludes []string) *GlobMatcher {
	return &GlobMatcher{
		Include: includes,
		Exclude: excludes,
	}
}

// Included checks if a given target path matches any of the include patterns.
// A matcher with no include patterns implicitly includes all targets.
func (m GlobMatcher) Included(target string) bool {
	if len(m.Include) == 0 {
		return true
	}
	return slices.Contains(m.Include, target) ||
		slices.ContainsFunc(m.Include, func(glob string) bool {
			ok, err := filepath.Match(glob, target)
			return err == nil && ok
		})
}

// Excluded checks if a given target path matches any of the exclude patterns.
func (m GlobMatcher) Excluded(target string) bool {
	return slices.Contains(m.Exclude, target) ||
		slices.ContainsFunc(m.Exclude, func(glob string) bool {
			ok, err := filepath.Match(glob, target)
			return err == nil && ok
		})
}

// Match determines if a target path is included and not excluded by the matcher's patterns.
func (m GlobMatcher) Match(target string) bool {
	return m.Included(target) && !m.Excluded(target)
}

// WalkFS traverses the file system starting from the root, applying the matcher's patterns to each file and directory.
// Matched files are processed by the provided walkFunc.
// The function accumulates errors encountered during traversal and processing, returning them as a single error.
func (m GlobMatcher) WalkFS(fsys fs.FS, walkFunc func(path string, content []byte) error) error {
	warnings := errs.Errors{}
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			warnings.WithErrorf("walk file step: %w", err)
			return nil
		}

		if d.IsDir() {
			if m.Excluded(path) {
				warnings.WithErrorf("skipping excluded dir '%s'", path)
				return fs.SkipDir
			}
			return nil
		}

		if m.Match(path) {
			content, err := fs.ReadFile(fsys, path)
			if err != nil {
				warnings.WithErrorf("read file: %w", err)
				return nil
			}

			if err := walkFunc(path, content); err != nil {
				warnings.WithErrorf("walk func file: %w", err)
			}
		}

		return nil
	})

	warnings.WithError(err)
	return warnings.OrNil()
}
