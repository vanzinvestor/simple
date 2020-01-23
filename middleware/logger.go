package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// customResponseWriter ...
type customResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

// newCustomResponseWriter ...
func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{
		ResponseWriter: w,
		StatusCode:     200,
	}
}

// Header ...
func (c *customResponseWriter) Header() http.Header {
	return c.ResponseWriter.Header()
}

// Write ...
func (c *customResponseWriter) Write(b []byte) (int, error) {
	return c.ResponseWriter.Write(b)
}

// WriteHeader ...
func (c *customResponseWriter) WriteHeader(statusCode int) {
	c.StatusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}

// Logger middleware logging response
func Logger(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		w2 := newCustomResponseWriter(w)

		next(w2, r, ps)

		duration := time.Now().Sub(start)
		log.Printf("%s %s%s %s %d %s", r.Method, r.Host, r.RequestURI, r.Proto, w2.StatusCode, duration)
	}
}
