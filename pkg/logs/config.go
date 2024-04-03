package logs

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"syscall"

	"github.com/toolvox/utilgo/pkg/errs"
	lh "github.com/toolvox/utilgo/pkg/logs/handlers"
	"github.com/toolvox/utilgo/pkg/timeutil"
)

// HandlerConfig defines the configuration for a log handler including its base behavior,
// target output, and additional options.
type HandlerConfig struct {
	Base    HandlerBase
	Target  HandlerTarget
	Options HandlerOption
}

// Handler constructs a new slog.Handler based on the HandlerConfig.
func (hc HandlerConfig) Handler() slog.Handler {
	if hc.Base == nil {
		hc.Base = JsonHandler{}
	}
	if hc.Target == nil {
		hc.Target = StderrTarget{}
	}
	if hc.Options == nil {
		hc.Options = make(HandlerOptions, 0)
	}

	handlerOptions := &slog.HandlerOptions{}
	hc.Options.SetOptions(handlerOptions)
	return hc.Base.GetHandler(hc.Target, handlerOptions)
}

// HandlerBase represents the base behavior required to get a slog.Handler.
type HandlerBase interface {
	GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler
}

// JsonHandler is a HandlerBase for creating JSON format log handlers.
type JsonHandler struct{}

func (base JsonHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(target.GetTarget(), options)
}

// TextHandler is a HandlerBase for creating plain text format log handlers.
type TextHandler struct{}

func (base TextHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(target.GetTarget(), options)
}

// CustomHandler allows for the use of a custom slog.Handler.
type CustomHandler struct {
	Handler slog.Handler
}

func (base CustomHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return base.Handler
}

// LogHandler wraps a utilgo's log handler into a HandlerBase.
type LogHandler struct {
	Log lh.Handler
}

func (base LogHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return base.Log.Handler()
}

// HandlerTarget defines the required method for a log output target.
type HandlerTarget interface {
	GetTarget() io.Writer
}

// StdoutTarget outputs logs to standard output.
type StdoutTarget struct{}

func (target StdoutTarget) GetTarget() io.Writer {
	return os.Stdout
}

// StderrTarget outputs logs to standard error.
type StderrTarget struct{}

func (target StderrTarget) GetTarget() io.Writer {
	return os.Stderr
}

// FileTarget outputs logs to a file, with optional timestamp prefixing.
type FileTarget struct {
	Name            string
	PrefixTimestamp bool
}

func (target FileTarget) GetTarget() io.Writer {
	dir, name := filepath.Dir(target.Name), filepath.Base(target.Name)
	if target.PrefixTimestamp {
		name = fmt.Sprintf("%s/%s_%s", dir, timeutil.TimestampNow(), name)
	}
	err := os.MkdirAll(dir, 0644)
	if err != nil && !errs.CheckPathError(err, "mkdir", dir, syscall.ERROR_ALREADY_EXISTS) {
		panic(err)
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

// WriterTarget wraps an io.Writer to be used as a log output target.
type WriterTarget struct {
	io.Writer
}

func (target WriterTarget) GetTarget() io.Writer {
	return target
}

// HandlerOption allows for configuration options to be set on a slog.Handler.
type HandlerOption interface {
	SetOptions(opt *slog.HandlerOptions)
}

// HandlerOptions is a list of HandlerOption(s) that get applied in order.
type HandlerOptions []HandlerOption

func (o HandlerOptions) SetOptions(opt *slog.HandlerOptions) {
	for _, option := range o {
		option.SetOptions(opt)
	}
}

// LogLevelOption sets the logging level for a handler.
type LogLevelOption slog.Level

func (o LogLevelOption) SetOptions(opt *slog.HandlerOptions) {
	opt.Level = slog.Level(o)
}

// AddSourceOption enables the addition of source information (like file and line number) to log entries.
type AddSourceOption struct{}

func (o AddSourceOption) SetOptions(opt *slog.HandlerOptions) {
	opt.AddSource = true
}

// ReplaceAttrOption allows for custom modification of log attributes.
type ReplaceAttrOption func(groups []string, a slog.Attr) slog.Attr

func (o ReplaceAttrOption) SetOptions(opt *slog.HandlerOptions) {
	if opt.ReplaceAttr == nil {
		opt.ReplaceAttr = o
		return
	}
	opt.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		a = opt.ReplaceAttr(groups, a)
		a = o(groups, a)
		return a
	}
}
