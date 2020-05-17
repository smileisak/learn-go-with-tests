package context

import (
	"errors"
	"net/http"
)

// SpyResponseWriter struct
type SpyResponseWriter struct {
	written bool
}

// Header func
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Write func
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

// WriteHeader func
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
