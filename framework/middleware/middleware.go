package middleware

import (
	"net/http"
)

type MiddlewareResponseWrite struct {
	http.ResponseWriter
	written bool
}

func NewMiddlewareResponseWriter(w http.ResponseWriter) *MiddlewareResponseWrite {
	return &MiddlewareResponseWrite{
		ResponseWriter: w,
	}
}

func (w *MiddlewareResponseWrite) Write(bytes []byte) (int, error) {
	w.written = true
	return w.ResponseWriter.Write(bytes)
}

func (w *MiddlewareResponseWrite) WriteHeader(code int) {
	w.written = true
	w.ResponseWriter.WriteHeader(code)
}

type Middleware []http.Handler

func (m *Middleware) Add(h http.Handler) {
	*m = append(*m, h)
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Wrap the supplied ResponseWriter
	mw := NewMiddlewareResponseWriter(w)

	// Loop through all of the registered handlers
	for _, handler := range m {
		// Call the handler with our MiddlewareResponseWriter
		handler.ServeHTTP(mw, r)

		// If there was a write, stop processing
		if mw.written {
			return
		}
	}
	// If no handlers wrote to the response, it’s a 404
	http.NotFound(w, r)
}
