package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {}

type WrappedResponseWriter struct{
	w http.ResponseWriter
	gz *gzip.Writer
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	gz := gzip.NewWriter(w)
	return &WrappedResponseWriter{w: w, gz: gz}
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		headerKey := r.Header.Get("Accept-Encoding")
		headerValue := "gzip"
		if strings.Contains(headerKey, headerValue){
			wrappedResponseWriter := NewWrappedResponseWriter(w)
			wrappedResponseWriter.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(wrappedResponseWriter, r)
			defer wrappedResponseWriter.Flush()
			return
		}
		next.ServeHTTP(w,r)
	})
}

func (wrappedResponseWriter *WrappedResponseWriter) Header() http.Header {
	return wrappedResponseWriter.w.Header()
}

func (wrappedResponseWriter *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wrappedResponseWriter.w.Write(d)
}

func (wrappedResponseWriter *WrappedResponseWriter) WriteHeader(statusCode int) {
	wrappedResponseWriter.w.WriteHeader(statusCode)
}

func (wrappedResponseWriter *WrappedResponseWriter) Flush() {
	wrappedResponseWriter.gz.Flush()
	wrappedResponseWriter.gz.Close()

}

