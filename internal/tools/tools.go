//go:build tools

package tools

import (
	//TODO2 remove codegen import after tool directive is added to mod.
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
	_ "github.com/twitchtv/circuitgen"
)
