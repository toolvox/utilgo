// Package logs provides a simplified interface for creating and configuring loggers with support for custom logging options.
// It leverages the slog package for structured logging and offers convenience functions to create loggers with various configurations.
package logs

import (
	"log/slog"

	lh "utilgo/pkg/logs/handlers"
)

// NewNullLogger creates a new [pkg/log/slog.Logger] that uses a [pkg/utilgo/pkg/logs/handlers.NullHandler], effectively discarding all logged messages.
// This is useful in contexts where logging is not required.
func NewNullLogger() *slog.Logger { return slog.New(lh.NullHandler{}) }

// NewLogger creates a new [pkg/log/slog.Logger] that combines multiple logging handlers into a [pkg/utilgo/pkg/logs/handlers.TeeHandler].
// This allows log messages to be dispatched to multiple handlers, each potentially with different logging configurations.
// if no handlers are provided, returns a [NewNullLogger].
func NewLogger(loggers ...LoggingOptions) *slog.Logger {
	var handlers []slog.Handler
	for _, logger := range loggers {
		handlers = append(handlers, logger.Handler())
	}
	return slog.New(lh.NewTeeHandler(handlers...))
}
