// Package handlers provides log handling mechanisms for various logging scenarios.
package handlers

import (
	"context"
	"log/slog"

	"utilgo/api"
)

// NullHandler is a no-op implementation of the [pkg/log/slog.Handler] interface that ignores any logging actions.
// It can be used when logs should be discarded or not handled.
type NullHandler api.Unit

// Enabled always returns false, indicating that this handler does not process any log level.
func (h NullHandler) Enabled(_ context.Context, _ slog.Level) bool { return false }

// Handle ignores the log record and returns nil, performing no action.
func (h NullHandler) Handle(_ context.Context, _ slog.Record) error { return nil }

// WithAttrs returns the same [NullHandler] instance, ignoring any attributes since it performs no action.
func (h NullHandler) WithAttrs(_ []slog.Attr) slog.Handler { return h }

// WithGroup returns the same [NullHandler] instance, ignoring any group specification since it performs no action.
func (h NullHandler) WithGroup(_ string) slog.Handler { return h }
