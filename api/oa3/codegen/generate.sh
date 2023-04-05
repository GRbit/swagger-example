#/bin/bash!

cd "$(dirname "$0")"
cd ../../..

go get github.com/deepmap/oapi-codegen/pkg/codegen@v1.12.4
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen | true

mkdir -p pkg/codegen/client
mkdir -p pkg/codegen/models
mkdir -p internal/codegen/handlers
mkdir -p internal/codegen/server

oapi-codegen -config api/oa3/codegen/server.yaml api/oa3/api.yaml
oapi-codegen -config api/oa3/codegen/client.yaml api/oa3/api.yaml
oapi-codegen -config api/oa3/codegen/models.yaml api/oa3/api.yaml

# copy yaml spec files to have them embed in binary
cp api/oa3/api.yaml internal/codegen/handlers/api.yaml
