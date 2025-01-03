package main

import (
	"context"
	"fmt"

	api "github.com/streemtech/twitchApi"
)

// Test exe to make sure the library is still compilable
func main() {

	c, err := api.NewClientWithResponses(api.SERVER, true)
	if err != nil {
		panic(err.Error())
	}

	resp, err := c.AddBlockedTermWithResponse(context.Background(), &api.AddBlockedTermParams{
		BroadcasterId: "29543631",
		ModeratorId:   "29543631",
	}, api.AddBlockedTermBody{
		Text: "foo",
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Status: %d", resp.StatusCode())
}
