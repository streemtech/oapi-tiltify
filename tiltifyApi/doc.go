package tiltifyApi

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config.types.yaml tiltify.openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config.server.yaml tiltify.openapi.yaml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=config.client.yaml tiltify.openapi.yaml
