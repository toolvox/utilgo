package logs

import "log/slog"

// Error creates a [pkg/log/slog.Attr] representing an error.
// If the error is nil, it creates an attribute with a value of "nil".
func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{
			Key:   "error",
			Value: slog.StringValue("nil"),
		}
	}
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
