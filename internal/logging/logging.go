package logging

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func InitLogger() error {
	l, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Logger = l.Sugar()
	return nil
}

// WithLogging is a middleware that logs HTTP requests and responses.
func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if Logger == nil {
			// Fallback to default logger or panic
			zap.NewExample().Sugar().Warn("Logger is not initialized, using fallback logger")
		}

		start := time.Now()
		uri := r.RequestURI
		method := r.Method
		ip := r.RemoteAddr
		bodySize := r.ContentLength

		responseData := &responseData{}
		lw := &loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		next.ServeHTTP(lw, r)

		duration := time.Since(start)

		// Use the initialized logger if available
		if Logger != nil {
			Logger.Infow(
				"Request processed",
				"uri", uri,
				"method", method,
				"duration", duration,
				"ip", ip,
				"body_size", bodySize,
				"status", responseData.status,
				"size", responseData.size,
			)
		}
	})
}

type responseData struct {
	status int
	size   int
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.responseData.size += size
	return size, err
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.ResponseWriter.WriteHeader(statusCode)
	lrw.responseData.status = statusCode
}
