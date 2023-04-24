// Package oapi_chi_server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package oapi_chi_server

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/grbit/swagger-example/pkg/codegen/models"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Retrieves properties information
	// (GET /property)
	GetPropertiesInfo(w http.ResponseWriter, r *http.Request)
	// Retrieves property information by ID
	// (GET /property/{propertyId})
	GetPropertyInfoById(w http.ResponseWriter, r *http.Request, propertyId string)
	// Set/Update property information by ID
	// (POST /property/{propertyId})
	CreateUpdatePropertyInfoById(w http.ResponseWriter, r *http.Request, propertyId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetPropertiesInfo operation middleware
func (siw *ServerInterfaceWrapper) GetPropertiesInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, OpenIdProdScopes, []string{""})

	ctx = context.WithValue(ctx, OpenIdSandboxScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPropertiesInfo(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPropertyInfoById operation middleware
func (siw *ServerInterfaceWrapper) GetPropertyInfoById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "propertyId" -------------
	var propertyId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "propertyId", runtime.ParamLocationPath, chi.URLParam(r, "propertyId"), &propertyId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "propertyId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, OpenIdProdScopes, []string{""})

	ctx = context.WithValue(ctx, OpenIdSandboxScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPropertyInfoById(w, r, propertyId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateUpdatePropertyInfoById operation middleware
func (siw *ServerInterfaceWrapper) CreateUpdatePropertyInfoById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "propertyId" -------------
	var propertyId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "propertyId", runtime.ParamLocationPath, chi.URLParam(r, "propertyId"), &propertyId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "propertyId", Err: err})
		return
	}

	ctx = context.WithValue(ctx, OpenIdProdScopes, []string{""})

	ctx = context.WithValue(ctx, OpenIdSandboxScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateUpdatePropertyInfoById(w, r, propertyId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/property", wrapper.GetPropertiesInfo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/property/{propertyId}", wrapper.GetPropertyInfoById)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/property/{propertyId}", wrapper.CreateUpdatePropertyInfoById)
	})

	return r
}

type GetPropertiesInfoRequestObject struct {
}

type GetPropertiesInfoResponseObject interface {
	VisitGetPropertiesInfoResponse(w http.ResponseWriter) error
}

type GetPropertiesInfo200JSONResponse PropertiesInfoResponse

func (response GetPropertiesInfo200JSONResponse) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertiesInfo400JSONResponse GenericErrorResponse

func (response GetPropertiesInfo400JSONResponse) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertiesInfo404Response struct {
}

func (response GetPropertiesInfo404Response) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetPropertiesInfo500JSONResponse GenericErrorResponse

func (response GetPropertiesInfo500JSONResponse) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertiesInfo501JSONResponse GenericErrorResponse

func (response GetPropertiesInfo501JSONResponse) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(501)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertiesInfodefaultResponse struct {
	StatusCode int
}

func (response GetPropertiesInfodefaultResponse) VisitGetPropertiesInfoResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type GetPropertyInfoByIdRequestObject struct {
	PropertyId string `json:"propertyId"`
}

type GetPropertyInfoByIdResponseObject interface {
	VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error
}

type GetPropertyInfoById200JSONResponse PropertyInfoResponse

func (response GetPropertyInfoById200JSONResponse) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertyInfoById400JSONResponse GenericErrorResponse

func (response GetPropertyInfoById400JSONResponse) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertyInfoById404Response struct {
}

func (response GetPropertyInfoById404Response) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetPropertyInfoById500JSONResponse GenericErrorResponse

func (response GetPropertyInfoById500JSONResponse) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertyInfoById501JSONResponse GenericErrorResponse

func (response GetPropertyInfoById501JSONResponse) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(501)

	return json.NewEncoder(w).Encode(response)
}

type GetPropertyInfoByIddefaultResponse struct {
	StatusCode int
}

func (response GetPropertyInfoByIddefaultResponse) VisitGetPropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

type CreateUpdatePropertyInfoByIdRequestObject struct {
	PropertyId string `json:"propertyId"`
	Body       *CreateUpdatePropertyInfoByIdJSONRequestBody
}

type CreateUpdatePropertyInfoByIdResponseObject interface {
	VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error
}

type CreateUpdatePropertyInfoById200JSONResponse PropertyInfoResponse

func (response CreateUpdatePropertyInfoById200JSONResponse) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateUpdatePropertyInfoById400JSONResponse GenericErrorResponse

func (response CreateUpdatePropertyInfoById400JSONResponse) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type CreateUpdatePropertyInfoById404Response struct {
}

func (response CreateUpdatePropertyInfoById404Response) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type CreateUpdatePropertyInfoById500JSONResponse GenericErrorResponse

func (response CreateUpdatePropertyInfoById500JSONResponse) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type CreateUpdatePropertyInfoById501JSONResponse GenericErrorResponse

func (response CreateUpdatePropertyInfoById501JSONResponse) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(501)

	return json.NewEncoder(w).Encode(response)
}

type CreateUpdatePropertyInfoByIddefaultResponse struct {
	StatusCode int
}

func (response CreateUpdatePropertyInfoByIddefaultResponse) VisitCreateUpdatePropertyInfoByIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(response.StatusCode)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Retrieves properties information
	// (GET /property)
	GetPropertiesInfo(ctx context.Context, request GetPropertiesInfoRequestObject) (GetPropertiesInfoResponseObject, error)
	// Retrieves property information by ID
	// (GET /property/{propertyId})
	GetPropertyInfoById(ctx context.Context, request GetPropertyInfoByIdRequestObject) (GetPropertyInfoByIdResponseObject, error)
	// Set/Update property information by ID
	// (POST /property/{propertyId})
	CreateUpdatePropertyInfoById(ctx context.Context, request CreateUpdatePropertyInfoByIdRequestObject) (CreateUpdatePropertyInfoByIdResponseObject, error)
}

type StrictHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetPropertiesInfo operation middleware
func (sh *strictHandler) GetPropertiesInfo(w http.ResponseWriter, r *http.Request) {
	var request GetPropertiesInfoRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPropertiesInfo(ctx, request.(GetPropertiesInfoRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPropertiesInfo")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPropertiesInfoResponseObject); ok {
		if err := validResponse.VisitGetPropertiesInfoResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// GetPropertyInfoById operation middleware
func (sh *strictHandler) GetPropertyInfoById(w http.ResponseWriter, r *http.Request, propertyId string) {
	var request GetPropertyInfoByIdRequestObject

	request.PropertyId = propertyId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetPropertyInfoById(ctx, request.(GetPropertyInfoByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPropertyInfoById")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetPropertyInfoByIdResponseObject); ok {
		if err := validResponse.VisitGetPropertyInfoByIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// CreateUpdatePropertyInfoById operation middleware
func (sh *strictHandler) CreateUpdatePropertyInfoById(w http.ResponseWriter, r *http.Request, propertyId string) {
	var request CreateUpdatePropertyInfoByIdRequestObject

	request.PropertyId = propertyId

	var body CreateUpdatePropertyInfoByIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUpdatePropertyInfoById(ctx, request.(CreateUpdatePropertyInfoByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUpdatePropertyInfoById")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(CreateUpdatePropertyInfoByIdResponseObject); ok {
		if err := validResponse.VisitCreateUpdatePropertyInfoByIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}