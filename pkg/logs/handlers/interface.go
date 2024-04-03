package handlers

import "log/slog"

type Handler interface {
	Handler() slog.Handler
}
