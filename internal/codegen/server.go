package codegen

import (
	"context"
	"errors"

	chi_server "github.com/grbit/swagger-example/internal/codegen/server/chi"
	echo_server "github.com/grbit/swagger-example/internal/codegen/server/echo"
)

var ErrNotImplemented = errors.New("not implemented")

type ChiServer struct{}

func (s ChiServer) GetPropertiesInfo(ctx context.Context, request chi_server.GetPropertiesInfoRequestObject) (chi_server.GetPropertiesInfoResponseObject, error) {
	return (chi_server.GetPropertiesInfoResponseObject)(nil), nil
}

func (s ChiServer) GetPropertyInfoById(
	ctx context.Context,
	request chi_server.GetPropertyInfoByIdRequestObject,
) (chi_server.GetPropertyInfoByIdResponseObject, error) {
	return (chi_server.GetPropertyInfoByIdResponseObject)(nil), nil
}

func (s ChiServer) CreateUpdatePropertyInfoById(
	ctx context.Context,
	request chi_server.CreateUpdatePropertyInfoByIdRequestObject,
) (chi_server.CreateUpdatePropertyInfoByIdResponseObject, error) {
	return (chi_server.CreateUpdatePropertyInfoByIdResponseObject)(nil), nil
}

type EchoServer struct{}

func (e EchoServer) GetPropertiesInfo(
	ctx context.Context,
	request echo_server.GetPropertiesInfoRequestObject,
) (echo_server.GetPropertiesInfoResponseObject, error) {
	return (echo_server.GetPropertiesInfoResponseObject)(nil), nil
}

func (e EchoServer) GetPropertyInfoById(
	ctx context.Context,
	request echo_server.GetPropertyInfoByIdRequestObject,
) (echo_server.GetPropertyInfoByIdResponseObject, error) {
	return (echo_server.GetPropertyInfoByIdResponseObject)(nil), nil
}

func (e EchoServer) CreateUpdatePropertyInfoById(
	ctx context.Context,
	request echo_server.CreateUpdatePropertyInfoByIdRequestObject,
) (echo_server.CreateUpdatePropertyInfoByIdResponseObject, error) {
	return (echo_server.CreateUpdatePropertyInfoByIdResponseObject)(nil), nil
}
