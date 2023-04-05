package codegen

import (
	"net/http"

	oapi_server "github.com/grbit/swagger-example/internal/codegen/server"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"golang.org/x/xerrors"
)

func NewRouter(lg zerolog.Logger) *chi.Mux {
	rootRouter := chi.NewRouter()
	rootRouter.Use( // first chain of middlewares to be executed
		ContextPrepareMiddleware(lg), // the first middleware to be executed
		RecoverMiddleware,
	)

	handler := oapi_server.NewStrictHandlerWithOptions(
		&Server{},
		[]oapi_server.StrictMiddlewareFunc{}, // third chain of middlewares to be executed
		oapi_server.StrictHTTPServerOptions{
			RequestErrorHandlerFunc:  RequestErrorHandlerFunc,
			ResponseErrorHandlerFunc: ResponseErrorHandlerFunc,
		},
	)

	oapi_server.HandlerWithOptions(handler, oapi_server.ChiServerOptions{
		BaseRouter: rootRouter,
		Middlewares: []oapi_server.MiddlewareFunc{ // second chain of middlewares to be executed
			RequestLoggingMiddleware,
		},
		ErrorHandlerFunc: RequestErrorHandlerFunc,
	})

	return rootRouter
}

func RequestErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	panic(xerrors.Errorf("swagger validation: %+v", err))
}

func ResponseErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	panic(err)
}
