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
