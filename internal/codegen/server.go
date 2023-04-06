package codegen

import (
	"context"
	"errors"

	oapi_server "github.com/grbit/swagger-example/internal/codegen/server"
)

type Server struct{}

var ErrNotImplemented = errors.New("not implemented")

func (s Server) GetPropertiesInfo(ctx context.Context, request oapi_server.GetPropertiesInfoRequestObject) (oapi_server.GetPropertiesInfoResponseObject, error) {
	return (oapi_server.GetPropertiesInfoResponseObject)(nil), nil
}

func (s Server) GetPropertyInfoById(
	ctx context.Context,
	request oapi_server.GetPropertyInfoByIdRequestObject,
) (oapi_server.GetPropertyInfoByIdResponseObject, error) {
	return (oapi_server.GetPropertyInfoByIdResponseObject)(nil), nil
}

func (s Server) CreateUpdatePropertyInfoById(
	ctx context.Context,
	request oapi_server.CreateUpdatePropertyInfoByIdRequestObject,
) (oapi_server.CreateUpdatePropertyInfoByIdResponseObject, error) {
	return (oapi_server.CreateUpdatePropertyInfoByIdResponseObject)(nil), nil
}
