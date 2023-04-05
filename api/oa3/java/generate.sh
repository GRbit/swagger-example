#!/bin/bash
# fail on error
set -e pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR

SWAGGERFILE=swagger-codegen-cli-3.0.34.jar
if [ -f "$SWAGGERFILE" ]; then
    echo "$SWAGGERFILE exists."
else
    echo "$SWAGGERFILE does not exist. Downloading it"
    wget https://repo1.maven.org/maven2/io/swagger/codegen/v3/swagger-codegen-cli/3.0.34/swagger-codegen-cli-3.0.34.jar
fi

docker build . -t nextpax_dist_api_swagger

cp ../api.yaml .
docker run -u $(id -u ${USER}):$(id -g ${USER}) --rm -v "`pwd`:/workspace" -w "/workspace" nextpax_dist_api_swagger ./generate-in-docker.sh
rm api.yaml

echo
echo "Applying post fixes:"
echo "Fixing omitempty value type..."
sed -i -e 's/ \([a-z0-9]\+\) `json:"\(.*\),omitempty"/ *\1 `json:"\2,omitempty"/g' output/go/models/*
echo "Fixing main package path..."
sed -i 's#sw "./go"#sw "github.com/grbit/swagger-example/internal/java"#g' output/main.go
echo "Fixing models package name..."
sed -i 's#package java#package models#g' output/go/models/*

echo
echo "--- copy generated files ---"
echo

mkdir -p ../../../internal/api
mkdir -p ../../../internal/java
mkdir -p ../../../pkg/java/models
mkdir -p ../../../cmd/java

cp output/api/swagger.yaml ../../../internal/api/java-generated.yaml
cp output/go/*.go ../../../internal/java
cp output/go/models/* ../../../pkg/java/models
cp output/main.go ../../../cmd/java/

rm -r output

echo
echo "***************************"
echo "***         DONE        ***"
echo "***************************"
echo
