// Package fsutil provides helpers and utilities for working with Files and FileSystems.
package fsutil

import (
	"io/fs"
	"path/filepath"
)

type SubDirFS struct {
	fs.FS
	SubDir string
}

func NewSubDirFS(src fs.FS, subDir string) *SubDirFS {
	return &SubDirFS{
		FS:     src,
		SubDir: subDir,
	}
}

func (s SubDirFS) Open(name string) (fs.File, error) {
	newPath := filepath.ToSlash(
		filepath.Join(s.SubDir, filepath.Clean(name)),
	)
	return s.FS.Open(newPath)
}
