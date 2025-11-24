

clean:
	find . -name '*.gen.go' -delete

fetch: 
	wget -N https://raw.githubusercontent.com/DmitryScaletta/twitch-api-swagger/refs/heads/main/openapi.yaml

gen: clean fetch generate
generate:
	GODEBUG=gotypesalias=0 go generate ./...

# Testbuild allows making sure the build is still valid to at-least some extent.
testbuild:
	cd cmd && go build -o main