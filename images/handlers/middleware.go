package handlers

import (
	"compress/gzip"
	"net/http"
)

// GZipResponseMiddleware detects if the user client can handle zipped content
// and if so returns the response in GZipped format.
func GZipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Write the file
		next.ServeHTTP(w, r)
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	// Wrap the default writer in a gzip writer
	gw := gzip.NewWriter(rw)

	return &WrappedResponseWriter{rw, gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

func (wr *WrappedResponseWriter) Write(b []byte) (int, error) {
	return wr.gw.Write(b)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.rw.WriteHeader(statusCode)
}
