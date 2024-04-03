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

type HandlerConfig struct {
	Base    HandlerBase
	Target  HandlerTarget
	Options *slog.HandlerOptions
}

func (hc HandlerConfig) Handler() slog.Handler {
	return hc.Base.GetHandler(hc.Target, hc.Options)
}

/**/

type HandlerBase interface {
	GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler
}

type JsonHandler struct{}

func (base JsonHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return slog.NewJSONHandler(target.GetTarget(), options)
}

type TextHandler struct{}

func (base TextHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return slog.NewTextHandler(target.GetTarget(), options)
}

type CustomHandler struct {
	Handler slog.Handler
}

func (base CustomHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return base.Handler
}

type LogHandler struct {
	Log lh.Handler
}

func (base LogHandler) GetHandler(target HandlerTarget, options *slog.HandlerOptions) slog.Handler {
	return base.Log.Handler()
}

/**/

type HandlerTarget interface{ GetTarget() io.Writer }

type StdoutTarget struct{}

func (target StdoutTarget) GetTarget() io.Writer {
	return os.Stdout
}

type StderrTarget struct{}

func (target StderrTarget) GetTarget() io.Writer {
	return os.Stderr
}

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

type WriterTarget struct {
	io.Writer
}

func (target WriterTarget) GetTarget() io.Writer {
	return target
}
