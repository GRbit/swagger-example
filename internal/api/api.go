package api

import "embed"

//go:embed java-generated.yaml
var swag embed.FS

// SwaggerSpec return bytes to serve swagger.yaml for https://github.com/swagger-api/swagger-ui
func SwaggerSpec() ([]byte, error) {
	return swag.ReadFile("java-generated.yaml")
}
