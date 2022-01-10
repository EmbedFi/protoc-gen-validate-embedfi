module github.com/embedfi/protoc-gen-embedfi

go 1.17

require (
	github.com/embedfi/finance v0.0.0-20220109050138-c7effce4d68c
	github.com/envoyproxy/protoc-gen-validate v0.4.0
	github.com/lyft/protoc-gen-star v0.6.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

// replace github.com/embedfi/finance => ../libraries/finance

replace github.com/envoyproxy/protoc-gen-validate => github.com/embedfi/protoc-gen-validate v0.4.2-0.20220109045928-0239dadd412b
