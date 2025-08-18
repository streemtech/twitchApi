# Twitch API

This is a golang twitch api client generated off of the [Twitch API Swagger](https://github.com/DmitryScaletta/twitch-api-swagger) by DmitryScaletta. It also generates a matching circuit wrapper via the [twitchtv/circuitgen](github.com/twitchtv/circuitgen) package.

## Correctness

This api is only as correct as the generated swagger files it is generating off of.

- If there is something missing that is in the swagger, please file a ticket [here](https://github.com/streemtech/twitchApi/issues/new) so that the api can be re-generated.
- If there is something incorrect or missing from the swagger, please file a ticket [here](https://github.com/DmitryScaletta/twitch-api-swagger/issues/new) so that the swagger can be corrected to generate properly

Under some circumstances, it may be necessary to file a ticket on the twitch api github issues.

Given this library may have lacking support for IE certain enums, or be incorrect in some circumstances due to how it is created from the documentation, You may wish to examine [nicklaw5/helix](https://github.com/nicklaw5/helix) which is hand-written, has few dependencies, and has strong tests for its endpoints if those are important to you.

## Updating

To update the api, simply re-run `make gen` to pull in the latest changes to the swagger file and update the library.

## License

This library is operating under an MIT license, nonwithstanding the license of github.com/DmitryScaletta/twitch-api-swagger (MIT license in the package.json file) and the license of the twitch api docs that the code is generated off of.

## Usage

The following is copied from examples/simple-circuit, but serves as a basic example on how to construct the client, its associated circuit, and make a request to the twitch api.

```go

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cep21/circuit/v3"
	"github.com/streemtech/twitchApi"
)

func main() {

	//Get the client ID and tokens from somewhere (ideally not as string constants)
	clientID := os.Getenv("TWITCH_CLIENT_ID")
	AppAccessToken := os.Getenv("TWITCH_APP_ACCESS_TOKEN")
	UserAccessToken := os.Getenv("TWITCH_INDIVIDUAL_USER_ACCESS_TOKEN")

	//load the eventsub client
	rootClient, err := twitchApi.NewClientWithResponses(
		twitchApi.SERVER, //Use the included server constant for the server host prefix
		true,             //Set this to true to manually handle non-200 responses. If false, any non-200 response will return an error with the response.
		twitchApi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			//Set the Client-Id and Authorization headers to be added automatically on each request
			req.Header.Set("Client-Id", clientID)
			//Set the app access token here if using
			req.Header.Set("Authorization", "Bearer "+AppAccessToken)
			return nil
		}),
	)
	if err != nil {
		log.Printf("failed to construct the twitch client ID: %s", err.Error())
		return
	}

	//optionally construct a circuit manager and wrap the client with it to automatically instrument all client calls.
	circuitManager := &circuit.Manager{
		//set the appropriate circuit manager configuration here as desired.
	}

	var circuitClient twitchApi.ClientWithResponsesInterface
	circuitClient, err = twitchApi.NewCircuitWrapperClientWithResponsesCircuit(circuitManager, rootClient, twitchApi.CircuitWrapperClientWithResponsesCircuitConfig{
		Defaults: circuit.Config{
			General: circuit.GeneralConfig{
				//Recommended default if using circuit only for metrics tracking.
				//Will mean excessive errors DO NOT cause the circuit to automatically error until request start resolving again.
				ForcedClosed: true,
			},
		},
        Prefix: "TwitchApiCircuit.",
	})
	if err != nil {
		log.Printf("failed to wrap client with circuit: %s", err.Error())
		return
	}

	//Use the client as desired
	resp, err := circuitClient.SendChatMessageWithResponse(context.Background(), twitchApi.SendChatMessageBody{
		BroadcasterId: "BROADCASTER",
		SenderId:      "SENDER",
		Message:       "Hi",
	}, func(ctx context.Context, req *http.Request) error {
		//If user access token is required, use set here on the individual request to override the authorization header set above.
		req.Header.Set("Authorization", "Bearer "+UserAccessToken)
		return nil
	})
	if err != nil {
		log.Printf("failed to send chat message to channel: %s", err.Error())
		return
	}
	log.Printf("Sent message, got ID %s", resp.JSON200.GetData()[0].MessageId)

}

```
