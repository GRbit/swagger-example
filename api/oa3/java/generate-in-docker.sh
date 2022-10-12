#!/bin/bash
# fail on error
set -e pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd $DIR

PATH=$PATH:$(go env GOPATH)/bin

printf "\n\n"
printf "***************************\n"
printf "***  Generating server  ***\n"
printf "***************************\n"
printf "\n"

java -jar swagger-codegen-cli-3.0.34.jar generate -i api.yaml \
 --additional-properties packageName=java,packageNameUpper=Java \
 -l go-server -t ./template -o ./output/

printf "\nApplying post fixes for models package name\n"
sed -i 's#package external#package models#g' output/go/model_*.go

mkdir ./output/go/models
mv ./output/go/model_*.go ./output/go/models

printf "\n"
printf "***************************\n"
printf "***  Generation DONE    ***\n"
printf "***************************\n"
printf "\n\n"
