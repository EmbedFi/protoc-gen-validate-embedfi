package main

import (
	"github.com/envoyproxy/protoc-gen-validate/module"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	scaled "github.com/embedfi/protoc-gen-validate-embedfi/scaled"
)

func main() {
	module.RegisterExtension(scaled.ScaledAmount{})
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV")).
		RegisterModule(module.Validator()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
