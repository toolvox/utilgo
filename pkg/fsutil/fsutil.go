// Package fsutil provides helpers and utilities for working with Files and FileSystems.
package fsutil

import (
	"io/fs"
	"path/filepath"
)

func ListFS(fsys fs.FS, root, glob string) ([]string, error) {
	var result []string
	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if ok, err := filepath.Match(glob, path); !ok || err != nil {
				return err
			}
			result = append(result, path)
		}
		return nil
	})
	return result, err
}
