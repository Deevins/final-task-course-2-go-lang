// Package docs содержит OpenAPI спецификацию для gateway.
package docs

import _ "embed"

//go:embed swagger.json
var swaggerJSON []byte

//go:embed swagger.yaml
var swaggerYAML []byte

// SwaggerJSON возвращает JSON спецификацию OpenAPI.
func SwaggerJSON() []byte {
	return swaggerJSON
}

// SwaggerYAML возвращает YAML спецификацию OpenAPI.
func SwaggerYAML() []byte {
	return swaggerYAML
}
