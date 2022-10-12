export TESTUSER=client_TST001:7fb3ec73-c5a7-4f2a-9c6a-8efeaeffa3d6

fmt:
	gofumpt -l -w -extra ./.
	gci write -s Std -s 'Prefix(github.com/grbit/swagger-example)' -s Def .

vet:
	go vet ./src/...

lint: fmt
	@golangci-lint run  --new-from-rev=master ./...

# https://github.com/swagger-api/swagger-codegen
oa3-java-generate:
	./api/oa3/java/generate.sh
	make fmt

# https://github.com/go-swagger/go-swagger
swagger2-generate:
	./api/sw2/generate.sh
	make fmt

start-swagger2:
	@go run ./cmd/stats-server/main.go
