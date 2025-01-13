//go:build tools

package tools

import (
	//TODO remove codegen import after tool directive is added to mod.
	_ "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/deepmap/oapi-codegen/v2/pkg/codegen"
	_ "github.com/rjeczalik/interfaces"
	_ "github.com/twitchtv/circuitgen"
	_ "github.com/vektra/mockery/v2"
)
