# for swagger2 server
export PORT?=8080

fmt:
	gofumpt -l -w -extra ./.
	gci write -s Std -s 'Prefix(github.com/grbit/swagger-example)' -s Def .

vet:
	go vet ./src/...

lint: fmt
	@golangci-lint run  --new-from-rev=master ./...

# https://github.com/swagger-api/swagger-codegen
java-generate:
	./api/oa3/java/generate.sh
	make fmt

# https://github.com/go-swagger/go-swagger
swagger2-generate:
	./api/sw2/generate.sh
	make fmt

# https://github.com/deepmap/oapi-codegen
codegen-generate:
	./api/sw2/generate.sh
	make fmt

start-swagger2:
	@go run ./cmd/stats-server/main.go

start-java:
	@go run ./cmd/java/main.go

start-codegen:
	@go run ./cmd/codegen/main.go
