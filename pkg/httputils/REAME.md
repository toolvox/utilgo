# httputil

* `requests.go` - basic http requests
  * `func JsonGet[TOut any](url string) (TOut, error) `
  * `func JsonPost[TIn, TOut any](url string, body TIn) (TOut, error) `
* `httputils.go`
  ```go
  type Mux interface {
    Handler(r *http.Request) (h http.Handler, pattern string)
    ServeHTTP(w http.ResponseWriter, r *http.Request)
    Handle(pattern string, handler http.Handler)
    HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
  }
  ```