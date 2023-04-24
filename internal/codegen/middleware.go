package codegen

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/segmentio/ksuid"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time) {
			zerolog.Ctx(r.Context()).Info().
				Dur("timeSpent", time.Since(t)).
				Msgf("%s request responded in %d ms", r.URL, time.Since(t).Milliseconds())
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zerolog.Ctx(r.Context()).Info().
			Str("requestURI", r.RequestURI).
			Str("httpMethod", r.Method).
			Str("remoteAddress", r.RemoteAddr).
			Msg("request received")

		next.ServeHTTP(w, r)
	})
}

func ContextPrepareMiddleware(lg zerolog.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			remoteAddr := r.Header.Get("X-Real-Ip")
			if remoteAddr == "" {
				remoteAddr = r.Header.Get("X-Forwarded-For")
			}
			if remoteAddr == "" {
				remoteAddr = r.RemoteAddr
			}

			lg = lg.With().
				Str("remoteAddress", remoteAddr).
				Str("httpMethod", r.Method).
				Str("requestID", ksuid.New().String()).
				Str("requestURI", r.RequestURI).
				Int64("contentLength", r.ContentLength).
				Logger()
			r = r.WithContext(lg.WithContext(r.Context()))

			next.ServeHTTP(w, r)
		})
	}
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			lg := zerolog.Ctx(r.Context())

			caughtError := recover()
			if caughtError == nil {
				return
			}

			if err, ok := caughtError.(error); ok {
				lg.Err(err).
					Str("stack", string(debug.Stack())).
					Str("detailedError", fmt.Sprintf("%+v", caughtError)).
					Msg("Recovering from panic")
			} else {
				lg.Error().
					Str("stack", string(debug.Stack())).
					Msg("Recovering from panic")
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "internal server error"}`))
		}()

		next.ServeHTTP(w, r)
	})
}

func ContextPrepareMiddlewareEcho(lg zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			remoteAddr := c.Request().Header.Get("X-Real-Ip")
			if remoteAddr == "" {
				remoteAddr = c.Request().Header.Get("X-Forwarded-For")
			}
			if remoteAddr == "" {
				remoteAddr = c.Request().RemoteAddr
			}

			lg = lg.With().
				Str("remoteAddress", remoteAddr).
				Str("httpMethod", c.Request().Method).
				Str("requestID", ksuid.New().String()).
				Str("requestURI", c.Request().RequestURI).
				Int64("contentLength", c.Request().ContentLength).
				Logger()
			c.Set("logger", lg)

			return next(c)
		}
	}
}

func RecoverMiddlewareEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			lg := c.Get("logger").(zerolog.Logger)

			caughtError := recover()
			if caughtError == nil {
				return
			}

			err, ok := caughtError.(error)
			if ok {
				lg.Err(err).
					Str("stack", string(debug.Stack())).
					Str("detailedError", fmt.Sprintf("%+v", caughtError)).
					Msg("Recovering from panic")
			} else {
				lg.Error().
					Str("stack", string(debug.Stack())).
					Msg("Recovering from panic")
			}

			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
		}()

		return next(c)
	}
}

func RequestLoggingMiddlewareEcho(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lg := c.Get("logger").(zerolog.Logger)
		lg.Info().
			Msg("request received")

		return next(c)
	}
}
