module github.com/embedfi/protoc-gen-validate-embedfi

go 1.17

require (
	github.com/embedfi/finance v0.0.0-20220112042104-bba783e58073
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/lyft/protoc-gen-star v0.6.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/spf13/afero v1.8.0 // indirect
	golang.org/x/text v0.3.7 // indirect
)

// replace github.com/embedfi/finance => ../libraries/finance

replace github.com/envoyproxy/protoc-gen-validate => github.com/embedfi/protoc-gen-validate v0.4.2-0.20220131003626-e3821fe5c6b0

//replace github.com/envoyproxy/protoc-gen-validate => ../protoc-gen-validate
