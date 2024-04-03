// 'countula' is a command-line tool designed to count lines of code in files across a directory tree, allowing for granular control over which files are included or excluded from the count.
//
// Flags:
//
//	-root:          Sets the root directory from which to start traversal.
//	                If not specified, 'countula' defaults to the current working directory.
//
//	-include:       Specifies a comma-separated list of glob patterns to include in the traversal.
//	                Only files matching at least one of these patterns are considered for counting.
//	                If no includes are provided, everything except explicitly excluded patterns is included.
//
//	-exclude:       Specifies a comma-separated list of glob patterns to exclude from the traversal.
//	                Any file matching at least one of these patterns is ignored.
//
//	-out:           Determines the destination for the output report.
//	                Users can specify a file path to direct the output to a file, or leave this flag unspecified to default the output to standard output.
//
//	-ignore-prefix: Specifies line prefixes that trigger skipping the line, useful for ignoring comments or specific code patterns.
//
//	-dir-mode:      When set, the output groups counts by directory.
//
//	-merge-mode:    When set, the output does not split counts by file extension, merging them into a single count per directory or globally.
//
// Example Usage:
//
//	$ countula -root "./source" -out "count_report.txt" -include "*.go,*.json,*.md" -exclude ".git" -ignore-prefix "//,#" -dir-mode
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/toolvox/utilgo/pkg/cli/flagutil"
	"github.com/toolvox/utilgo/pkg/cmdutil"
	"github.com/toolvox/utilgo/pkg/errs"
	"github.com/toolvox/utilgo/pkg/fsutil"
	"github.com/toolvox/utilgo/pkg/logs"
	"github.com/toolvox/utilgo/pkg/maputil"
	"github.com/toolvox/utilgo/pkg/reflectutil"
)

const Version = "v0.1.1"

type countulaOpts struct {
	cmdutil.TraverseFlags

	IgnoreLinePrefix flagutil.CSVValue
	OutputFile       flagutil.OutputFileValue
	DirMode          bool
	MergeMode        bool

	Debug bool
}

func main() {
	var o countulaOpts
	o.TraverseFlags.RegisterVars(flag.CommandLine)
	flag.Var(&o.OutputFile, "out", "Output file path. Defaults to stdout if empty")
	flag.Var(&o.IgnoreLinePrefix, "ignore-prefix", "line prefixes that trigger skipping the line (comma separated)")
	flag.BoolVar(&o.DirMode, "dir-mode", false, "Group result by sub-directory")
	flag.BoolVar(&o.MergeMode, "merge-mode", false, "Don't split by extension")
	flag.BoolVar(&o.Debug, "debug", false, "print debug logs")

	flag.Parse()

	var level slog.Level = slog.LevelInfo
	if o.Debug {
		level = slog.LevelDebug
	}

	log := logs.NewLogger(
		logs.HandlerConfig{Options: logs.LogLevelOption(level)},
	)
	log.Info("started countula", slog.String("version", Version))
	if err := run(o, log); err != nil {
		log.Error("codump failed", logs.Error(err))
		os.Exit(1)
	}
}

func run(o countulaOpts, log *slog.Logger) error {
	log.Debug("get output writer for", slog.String("output file", o.OutputFile.Filename))
	writer := o.OutputFile.Writer()
	if writer == nil {
		return errs.New("error obtaining output writer")
	}
	defer writer.Close()

	log.Debug("isolating target fs")
	fsys := o.RootFS()

	log.Debug("loading ignores")
	ignores, err := o.Ignores(fsys)
	if err != nil {
		log.Warn("could not read ignore file", logs.Error(err))
	}

	log.Debug("includes/exclude from flags")

	includes := o.IncludeGlobs.Values
	log.Info("include globs", slog.Any("patterns", includes))
	excludes := append(o.ExcludeGlobs.Values, ignores...)
	excludes = append(excludes, o.OutputFile.String())
	log.Info("excluding globs", slog.Any("patterns", excludes))

	log.Debug("create matcher")
	matcher := fsutil.NewGlobMatcher(includes, excludes)
	skipPrefixes := o.IgnoreLinePrefix.Values
	counts := map[string]map[string]int{}

	log.Debug("count lines")
	matcher.WalkFS(fsys, func(path string, content []byte) error {
		log.Debug("scanning file", slog.String("path", path))
		ext := filepath.Ext(path)
		if ext == "" {
			ext = filepath.Base(path)
		}
		subDir := filepath.ToSlash(filepath.Dir(path))
		subDirCounts, ok := counts[subDir]
		if !ok {
			subDirCounts = make(map[string]int)
			counts[subDir] = subDirCounts
		}

		lines := strings.Split(string(content), "\n")
		lines = slices.DeleteFunc(lines, reflectutil.IsZero)
		lines = slices.DeleteFunc(lines, func(s string) bool {
			s = strings.TrimSpace(s)
			if len(s) <= 3 {
				return true
			}
			for _, prefix := range skipPrefixes {
				if strings.HasPrefix(s, prefix) {
					return true
				}
			}
			return false
		})

		subDirCounts[ext] += len(lines)
		return nil
	})

	log.Debug("print summary")
	writer.Write([]byte(renderCounts(counts, o.DirMode, o.MergeMode)))

	return nil
}

func renderCounts(counts map[string]map[string]int, dirMode bool, mergeMode bool) string {
	var sb strings.Builder
	total := 0

	if dirMode {
		dirKeys := maputil.SortedKeys(counts)
		for i, dir := range dirKeys {
			extCounts := counts[dir]
			if i > 0 {
				sb.WriteRune('\n')
			}
			sb.WriteString(fmt.Sprintf("Directory `%s`:\n", dir))
			exts := maputil.SortedKeys(extCounts)
			dirCount := 0
			for _, ext := range exts {
				count := extCounts[ext]
				total += count
				dirCount += count
				if !mergeMode {
					sb.WriteString(fmt.Sprintf("    `%s`: %d\n", ext, count))
				}
			}
			if mergeMode {
				sb.WriteString(fmt.Sprintf("    lines: %d\n", dirCount))
			}
		}
		sb.WriteString(fmt.Sprintf("\n  ===  \nTotal lines: %d\n", total))
		return sb.String()
	}

	if mergeMode {
		dirKeys := maputil.SortedKeys(counts)
		for _, dir := range dirKeys {
			extCounts := counts[dir]
			for _, count := range extCounts {
				total += count
			}
		}
		sb.WriteString(fmt.Sprintf("Total lines: %d\n", total))
		return sb.String()
	}

	extCounts := make(map[string]int)
	dirKeys := maputil.SortedKeys(counts)
	for _, dir := range dirKeys {
		dirCounts := counts[dir]

		exts := maputil.SortedKeys(dirCounts)
		for _, ext := range exts {
			count := dirCounts[ext]
			extCounts[ext] += count
			total += count
		}
	}

	sb.WriteString("Counts by extension:\n")
	exts := maputil.SortedKeys(extCounts)
	for _, ext := range exts {
		sb.WriteString(fmt.Sprintf("    `%s`: %d\n", ext, extCounts[ext]))
	}

	sb.WriteString(fmt.Sprintf("\n  ===  \nTotal lines: %d\n", total))
	return sb.String()
}
