package logs

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"utilgo/pkg/timeutil"
)

// LoggingOptions defines options for configuring log handlers.
// This includes output targets, whether to prefix file names with timestamps, logging formats, and levels.
// It also supports specifying custom handlers for complete control over logging behavior.
type LoggingOptions struct {
	// Filepath | "", "2", or "stderr" for stderr | "1", or "stdout" for stdout (default "")
	//
	// Overridden by [LoggingOptions.CustomHandler]
	Target string
	// adds a "2006_01_02_15_04_05_" timestamp prefix to the log file name
	// Not compatible with stderr/stdout (default false)
	PrefixFilenameWithTime bool
	// Whether to add the source location of the call to log (default false)
	AddSource bool
	// Minimum level to log for this handler (default INFO)
	Level slog.Level
	// true means TextHandler, false means JsonHandler (default false).
	//
	// Overridden by [LoggingOptions.CustomHandler]
	TextHandler bool
	// ReplaceAttrs are called to rewrite each attribute before it is logged. (default nil)
	ReplaceAttrs []func(attrGroups []string, a slog.Attr) slog.Attr
	// Your custom handler, overrides all other fields (default nil)
	CustomHandler slog.Handler
}

// Handler constructs a [pkg/log/slog.Handler] based on the [LoggingOptions].
// It supports creating file-based handlers, stdout/stderr handlers, and custom handlers.
func (o LoggingOptions) Handler() slog.Handler {
	if o.CustomHandler != nil {
		return o.CustomHandler
	}
	var file *os.File
	switch strings.ToLower(o.Target) {
	case "", "stderr":
		file = os.Stderr
	case "stdout":
		file = os.Stdout
	default:
		dir, name := filepath.Dir(o.Target), filepath.Base(o.Target)
		if o.PrefixFilenameWithTime {
			o.Target = fmt.Sprintf("%s/%s_%s", dir, timeutil.TimestampNow(), name)
		}
		os.MkdirAll(dir, 0644)
		var err error
		file, err = os.OpenFile(o.Target, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	}
	handlerOptions := &slog.HandlerOptions{
		AddSource: o.AddSource,
		Level:     o.Level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			for _, rep := range o.ReplaceAttrs {
				a = rep(groups, a)
			}
			return a
		},
	}
	if !o.TextHandler {
		return slog.NewJSONHandler(file, handlerOptions)
	} else {
		return slog.NewTextHandler(file, handlerOptions)
	}
}
