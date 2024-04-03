// 'codump' is a command-line utility designed to traverse a directory tree and "compile" the contents of files that match specified criteria into a single output file or stream.
// It leverages flexible include and exclude filtering capabilities, allowing users to specify precisely which files should be aggregated based on glob patterns.
// The utility supports output to both files and standard output (/error), making it versatile for various scripting and logging purposes.
//
// Flags:
//
//	-root:    Sets the root directory from which to start traversal.
//	          If not specified, 'codump' defaults to the current working directory.
//
//	-include: Specifies a comma-separated list of glob patterns to include in the traversal.
//	          Only files matching at least one of these patterns are considered for content aggregation.
//	          If no includes are provided, include everything that isn't excluded.
//
//	-exclude: Specifies a comma-separated list of glob patterns to exclude from the traversal.
//	          Any file matching at least one of these patterns is ignored.
//
//	-out:     Determines the destination for the aggregated content.
//	          Users can specify a file path to direct the output to a file or leave this flag unspecified to default the output to standard output.
//
// Example Usage:
//
//	$ codump -root "./project/target" -ignore ".gitignore" -exclude ".git" -include "*.go,*.md"
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/toolvox/utilgo/pkg/cli/flagutil"
	"github.com/toolvox/utilgo/pkg/cmdutil"
	"github.com/toolvox/utilgo/pkg/errs"
	"github.com/toolvox/utilgo/pkg/fsutil"
	"github.com/toolvox/utilgo/pkg/logs"
)

const Version = "v0.1.0"

type codumpOpts struct {
	cmdutil.TraverseFlags

	OutputFile flagutil.OutputFileValue
}

func main() {
	var o codumpOpts
	o.TraverseFlags.RegisterVars(flag.CommandLine)
	flag.Var(&o.OutputFile, "out", "Output file path. Defaults to stdout if empty")
	flag.Parse()

	log := logs.NewLogger(logs.LoggingOptions{Level: slog.LevelDebug})
	log.Info("started codump", slog.String("version", Version))
	if err := run(o, log); err != nil {
		log.Error("codump failed", logs.Error(err))
		os.Exit(1)
	}
}

func run(o codumpOpts, log *slog.Logger) error {
	// 1. Get Output Writer
	writer := o.OutputFile.Writer()
	if writer == nil {
		return errs.New("error obtaining output writer")
	}
	defer writer.Close()

	// 2. Isolate the target FS
	fsys := o.RootFS()

	// 3. Attempt to load the ignores
	ignores, err := o.Ignores(fsys)
	if err != nil {
		log.Warn("could not read ignore file", logs.Error(err))
	}

	// 4. Load includes/excludes
	includes := o.IncludeGlobs.Values
	log.Info("include globs", slog.Any("patterns", includes))
	excludes := append(o.ExcludeGlobs.Values, ignores...)
	excludes = append(excludes, o.OutputFile.String())
	log.Info("excluding globs", slog.Any("patterns", excludes))

	// 5. Create Matcher
	matcher := fsutil.NewGlobMatcher(includes, excludes)

	// 6. Run Matcher and output to writer
	warns := matcher.WalkFS(fsys, func(path string, content []byte) error {
		fmt.Fprintln(writer, "---", path)
		fmt.Fprintf(writer, "%s\n---\n", content)
		return nil
	})

	// 7. Log walk errors
	if warns != nil {
		log.Warn("walk dir warnings", logs.Error(warns))
	}

	return nil
}
