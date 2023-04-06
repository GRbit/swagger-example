// This file is safe to edit. Once it exists it will not be overwritten

package sw2

import (
	"crypto/tls"
	"net/http"

	"github.com/grbit/swagger-example/internal/sw2/operations"
	"github.com/grbit/swagger-example/pkg/sw2/model"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

//go:generate swagger generate server --target ../../../swagger-example --name Stats --spec ../../api/sw2/api.yaml --model-package pkg/sw2/model --server-package internal/sw2 --principal model.Principal

func configureFlags(api *operations.StatsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StatsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-System-Token" header is set
	api.APIKeyAuthAuth = func(token string) (*model.Principal, error) {
		// will authorize everything
		return nil, nil
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AddAccUpdaterHandler == nil {
		api.AddAccUpdaterHandler = operations.AddAccUpdaterHandlerFunc(func(params operations.AddAccUpdaterParams, principal *model.Principal) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddAccUpdater has not yet been implemented")
		})
	}
	if api.DelAccUpdaterHandler == nil {
		api.DelAccUpdaterHandler = operations.DelAccUpdaterHandlerFunc(func(params operations.DelAccUpdaterParams, principal *model.Principal) middleware.Responder {
			return middleware.NotImplemented("operation operations.DelAccUpdater has not yet been implemented")
		})
	}
	if api.GetAdsStatsHandler == nil {
		api.GetAdsStatsHandler = operations.GetAdsStatsHandlerFunc(func(params operations.GetAdsStatsParams, principal *model.Principal) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetAdsStats has not yet been implemented")
		})
	}
	if api.GetCampaignsStatsHandler == nil {
		api.GetCampaignsStatsHandler = operations.GetCampaignsStatsHandlerFunc(func(params operations.GetCampaignsStatsParams, principal *model.Principal) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetCampaignsStats has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
