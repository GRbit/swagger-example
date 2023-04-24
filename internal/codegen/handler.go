package codegen

import (
	"net/http"

	oapi_chi_server "github.com/grbit/swagger-example/internal/codegen/server/chi"
	oapi_echo_server "github.com/grbit/swagger-example/internal/codegen/server/echo"

	"github.com/go-chi/chi/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

func NewChiHandler(lg zerolog.Logger) http.Handler {
	rootRouter := chi.NewRouter()
	rootRouter.Use( // first chain of middlewares to be executed
		ContextPrepareMiddleware(lg), // the first middleware to be executed
		RecoverMiddleware,
	)

	handler := oapi_chi_server.NewStrictHandlerWithOptions(
		&ChiServer{},
		[]oapi_chi_server.StrictMiddlewareFunc{}, // third chain of middlewares to be executed
		oapi_chi_server.StrictHTTPServerOptions{
			RequestErrorHandlerFunc:  RequestErrorHandlerFunc,
			ResponseErrorHandlerFunc: ResponseErrorHandlerFunc,
		},
	)

	oapi_chi_server.HandlerWithOptions(handler, oapi_chi_server.ChiServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: rootRouter,
		Middlewares: []oapi_chi_server.MiddlewareFunc{ // second chain of middlewares to be executed
			RequestLoggingMiddleware,
		},
		ErrorHandlerFunc: RequestErrorHandlerFunc,
	})

	return rootRouter
}

func RequestErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
}

func ResponseErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	panic(xerrors.Errorf("response error: %w", err))
}

func NewEchoHandler(lg zerolog.Logger) http.Handler {
	rootRouter := echo.New()
	rootRouter.Use( // first chain of middlewares to be executed
		ContextPrepareMiddlewareEcho(lg), // the first middleware to be executed
		RecoverMiddlewareEcho,
		RecoverMiddlewareEcho,
	)

	serverInterface := oapi_echo_server.NewStrictHandler(
		&EchoServer{},
		[]oapi_echo_server.StrictMiddlewareFunc{}, // third chain of middlewares to be executed
	)

	oapi_echo_server.RegisterHandlersWithBaseURL(
		rootRouter,
		serverInterface,
		"/api/v1",
	)

	return rootRouter
}
