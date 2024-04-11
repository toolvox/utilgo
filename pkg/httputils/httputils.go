package httputils

import "net/http"

type Mux interface {
	Handler(r *http.Request) (h http.Handler, pattern string)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Handle(pattern string, handler http.Handler)
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type Middleware interface {
	Middleware(http.Handler) http.Handler
}

type MiddlewareFunc func(http.Handler) http.Handler

func (mf MiddlewareFunc) Middleware(next http.Handler) http.Handler {
	return mf(next)
}

type MiddlewareHandler interface {
	http.Handler

	Use(...MiddlewareFunc)
}

type ServerHandler struct{ http.Handler }

func (s *ServerHandler) Use(mws ...MiddlewareFunc) {
	for _, mw := range mws {
		s.Handler = mw.Middleware(s.Handler)
	}
}

func HandleMux(mux Mux) *ServerHandler {
	return &ServerHandler{Handler: mux}
}
