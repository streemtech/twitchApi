package api

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml https://raw.githubusercontent.com/DmitryScaletta/twitch-api-swagger/refs/heads/main/openapi.yaml
//generate API code.

//go:generate go run github.com/twitchtv/circuitgen --name ClientWithResponsesInterface --pkg ./ --debug --out ./circuit.gen.go --alias ClientWithResponsesCircuit --circuit-major-version 3
//generate circuits for interface
