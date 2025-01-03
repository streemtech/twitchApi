# Twitch API

This is a golang twitch api client generated off of the [Twitch API Swagger](https://github.com/DmitryScaletta/twitch-api-swagger) by DmitryScaletta.

## Correctness

This api is only as correct as the generated swagger files it is generating off of.

- If there is something missing, please file a ticket [here](https://github.com/streemtech/twitchApi/issues/new) so that the api can be re-generated.
- If there is something incorrect, please file a ticket [here](https://github.com/DmitryScaletta/twitch-api-swagger/issues/new) so that the swagger can be corrected to generate properly

## Updating

To update the api, simply re-run `make gen` to pull in the latest changes to the swagger file and update the library.

## License

This library is operating under an MIT license, nonwithstanding the license of github.com/DmitryScaletta/twitch-api-swagger and the license of the twitch api docs that the code is generated off of.
