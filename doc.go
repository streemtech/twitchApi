package twitchApi

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=config.yaml openapi.yaml
//generate API code.

//go:generate go run github.com/vektra/mockery/v2@v2.52.2 --name ClientWithResponsesInterface --output ./mocks --with-expecter --filename api.gen.go --structname Api
//generate client interface mocks

//go:generate go run github.com/twitchtv/circuitgen --name ClientWithResponsesInterface --pkg ./ --debug --out ./circuit.gen.go --alias ClientWithResponsesCircuit --circuit-major-version 3
//generate circuits for interface

// SERVER is the default helix server api url. Use with the creation of the client.
const SERVER = "https://api.twitch.tv/helix"
