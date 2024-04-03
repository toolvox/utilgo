package handlers

import (
	"context"
	"log/slog"

	"github.com/toolvox/utilgo/pkg/errs"
)

// TeeHandler is a composite log handler that forwards log records to a slice of [pkg/log/slog.Handler].
// It allows logs to be processed by multiple handlers simultaneously.
type TeeHandler []slog.Handler

// NewTeeHandler creates a new [TeeHandler] with the provided [pkg/log/slog.Handler] slice.
// If no handlers are provided, it returns a [NullHandler] that ignores all log actions.
// If only one handler is provided, it returns that handler.
func NewTeeHandler(handlers ...slog.Handler) slog.Handler {
	switch len(handlers) {
	case 0:
		return NullHandler{}
	case 1:
		return handlers[0]
	default:
		return TeeHandler(handlers)
	}
}

// Enabled checks if at least one of the contained handlers is enabled for the given log level and context.
// It returns true if any handler is enabled, otherwise false.
func (h TeeHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	for _, handler := range h {
		if handler.Enabled(ctx, lvl) {
			return true
		}
	}
	return false
}

// Handle forwards the log record to all enabled handlers within the [TeeHandler].
// If any handler returns an error, those errors are collected and returned as a single error.
func (h TeeHandler) Handle(ctx context.Context, record slog.Record) error {
	var errors errs.Errors
	for _, handler := range h {
		if !handler.Enabled(ctx, record.Level) {
			continue
		}
		if err := handler.Handle(ctx, record); err != nil {
			errors.WithError(err)
		}
	}
	return errors.OrNil()
}

// WithAttrs returns a new [TeeHandler] instance where each contained handler has been augmented with the provided attributes.
// It allows for attribute modifications to propagate through all handlers in the [TeeHandler].
func (h TeeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	var newHandlers []slog.Handler = make([]slog.Handler, len(h))
	for i, handler := range h {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return NewTeeHandler(newHandlers...)
}

// WithGroup returns a new [TeeHandler] instance where each contained handler is part of the specified group.
// This method allows for logical grouping of handlers within the [TeeHandler].
func (h TeeHandler) WithGroup(name string) slog.Handler {
	var newHandlers []slog.Handler = make([]slog.Handler, len(h))
	for i, handler := range h {
		newHandlers[i] = handler.WithGroup(name)
	}
	return NewTeeHandler(newHandlers...)
}
